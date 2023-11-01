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
	userHandler *handler.UserHandler,
) *gin.Engine {
	if conf.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(gin.Logger(), recovery.Handler())

	// cors config
	//router.Use(corsM.Handler())

	//apiGroup := router.Group("/api")
	// account handler
	//registerAccountHandler(apiGroup.Group("/account"), codeLimiter, localAuthLimiter, accountHandler)

	return router
}

//func registerTalkHandler(group *gin.RouterGroup, auth *middleware.Auth, handler *handler.TalkHandler) {
//	groupRouter := group.Use(auth.Handler())
//	groupRouter.POST("/audio/token", handler.CreateAudioRecognitionToken)
//	groupRouter.POST("/advice", handler.CreateTalkAdvice)
//	groupRouter.POST("/reply", handler.CreateTalkReply)
//	groupRouter.POST("/upload/token", handler.CreateUploadToken)
//	groupRouter.POST("/translate", handler.CreateTranslate)
//}
