package main

import (
	"ShortURL/internal/config"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	conf    *config.Configuration
	logger  *zap.Logger
	httpSrv *http.Server
}

func newHttpServer(
	conf *config.Configuration,
	router *gin.Engine,
) *http.Server {
	return &http.Server{
		Addr:    ":" + conf.App.Port,
		Handler: router,
	}
}

func newApp(
	conf *config.Configuration,
	logger *zap.Logger,
	httpSrv *http.Server,
) *App {
	return &App{
		conf:    conf,
		logger:  logger,
		httpSrv: httpSrv,
	}
}

func (a *App) Run() error {
	// action before start app
	go func() {
		a.logger.Info("app init")
	}()
	// start http server
	go func() {
		a.logger.Info("http server started")
		if err := a.httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return nil
}

func (a *App) Stop(ctx context.Context) error {
	// stop http server
	a.logger.Info("http server has been stop")
	if err := a.httpSrv.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

func (a *App) AwaitSignal() {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	select {
	case s := <-c:
		a.logger.Info("receive a signal", zap.String("signal", s.String()))

		// set 10 seconds timeout
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		_ = a.Stop(ctx)

		os.Exit(0)
	}
}
