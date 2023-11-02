package router

import (
	"ShortURL/internal/api/handler"
	"ShortURL/internal/api/middleware"
	"ShortURL/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// ProviderSet is router providers.
var ProviderSet = wire.NewSet(NewRouter)

func NewRouter(
	conf *config.Configuration,
	recovery *middleware.Recovery,
	idHandler *handler.IdHandler,
	urlHandler *handler.UrlHandler,
) *gin.Engine {
	if conf.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(gin.Logger(), recovery.Handler())

	// cors config
	//router.Use(corsM.Handler())

	apiGroup := router.Group("/api")
	//apiGroup.POST("/id", idHandler.NextId) // only for test
	apiGroup.POST("/url", urlHandler.ShortenUrl)

	shortenGroup := router.Group("/s")
	shortenGroup.GET("/:id", urlHandler.ExpandUrl)
	return router
}
