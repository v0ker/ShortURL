package main

import (
	"ShortURL/internal/config"
	"ShortURL/internal/utils"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	Version      string
	configPath   string
	conf         *config.Configuration
	loggerWriter *lumberjack.Logger
	logger       *zap.Logger
)

func init() {
	flag.StringVar(&configPath, "config", "../../configs", "config path, eg: -config config.yaml")
}

func main() {
	flag.Parse()

	initConfig()
	initLogger()

	app, cleanup, err := wireApp(conf, loggerWriter, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start app
	log.Printf("start app %s ...", Version)
	if err := app.Run(); err != nil {
		panic(err)
	}

	app.AwaitSignal()
}

func initConfig() {
	fmt.Println("load config:" + configPath)

	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}

	if err := v.Unmarshal(&conf); err != nil {
		fmt.Println(err)
	}
	bytes, _ := json.Marshal(conf)
	fmt.Printf("config file loaded: %s\n", string(bytes))
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		defer func() {
			if err := recover(); err != nil {
				logger.Error("config file changed err:", zap.Any("err", err))
				fmt.Println(err)
			}
		}()
		if err := v.Unmarshal(&conf); err != nil {
			fmt.Println(err)
		}
	})
}

func initLogger() {
	var level zapcore.Level  // zap log level
	var options []zap.Option // zap log options

	logFileDir := conf.Log.LogDir
	fmt.Printf("log file dir: %s\n", logFileDir)
	if ok, _ := utils.Exists(logFileDir); !ok {
		_ = os.Mkdir(logFileDir, os.ModePerm)
	}

	switch conf.Log.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.DebugLevel
	}

	// change encoder config
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("2006-01-02 15:04:05.000"))
	}
	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(conf.App.Env + "." + l.String())
	}

	loggerWriter = &lumberjack.Logger{
		Filename:   filepath.Join(logFileDir, conf.Log.Filename),
		MaxSize:    conf.Log.MaxSize,
		MaxBackups: conf.Log.MaxBackups,
		MaxAge:     conf.Log.MaxAge,
		Compress:   conf.Log.Compress,
	}
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(loggerWriter), level),
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(os.Stdout), level),
	)

	logger = zap.New(core, options...)
}
