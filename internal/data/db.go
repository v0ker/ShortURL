package data

import (
	"ShortURL/internal/config"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
	"time"
)

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

// NewData .
func NewData(c *config.Configuration, logger *zap.Logger) (*Data, func(), error) {
	db, dbCleanup, err := genGormDb(c, logger)
	if err != nil {
		panic("failed to connect database")
	}
	rdb, rdbCleanup, err := genRedis(c, logger)
	if err != nil {
		panic("failed to connect redis")
	}
	cleanup := func() {
		dbCleanup()
		rdbCleanup()
	}
	return &Data{
		db:  db,
		rdb: rdb,
	}, cleanup, nil
}

func genGormDb(c *config.Configuration, logger *zap.Logger) (*gorm.DB, func(), error) {
	cleanup := func() {
		logger.Info("closing the data resources")
	}
	dsn := c.Database.Source
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: glogger.Default.LogMode(glogger.Info),
	})
	if err != nil {
		return nil, cleanup, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, cleanup, err
	}
	sqlDB.SetMaxIdleConns(c.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Second * 60)
	return db, cleanup, err
}

func genRedis(c *config.Configuration, logger *zap.Logger) (*redis.Client, func(), error) {
	cleanup := func() {
		logger.Info("closing the redis resources")
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Address,
		Password: c.Redis.Password, // no password set
	})
	return rdb, cleanup, nil
}
