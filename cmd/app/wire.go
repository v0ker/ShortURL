//go:build wireinject
// +build wireinject

package main

import (
	"ShortURL/internal/api/handler"
	"ShortURL/internal/api/middleware"
	"ShortURL/internal/api/router"
	"ShortURL/internal/biz"
	"ShortURL/internal/config"
	"ShortURL/internal/data"
	"github.com/google/wire"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
)

// wireApp init application.
func wireApp(*config.Configuration, *lumberjack.Logger, *zap.Logger) (*App, func(), error) {
	panic(
		wire.Build(
			handler.ProviderSet,
			router.ProviderSet,
			middleware.ProviderSet,
			biz.ProviderSet,
			data.ProviderSet,
			//http and app
			newHttpServer,
			newApp,
		),
	)
}
