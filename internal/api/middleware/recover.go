package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
)

type Recovery struct {
	log          *zap.Logger
	loggerWriter *lumberjack.Logger
}

func NewRecovery(loggerWriter *lumberjack.Logger, logger *zap.Logger) *Recovery {
	return &Recovery{
		log:          logger,
		loggerWriter: loggerWriter,
	}
}

func (m *Recovery) Handler() gin.HandlerFunc {
	return gin.RecoveryWithWriter(
		m.loggerWriter,
		m.ServerError,
	)
}

func (m *Recovery) ServerError(c *gin.Context, err interface{}) {
	m.log.Error("server error", zap.Any("error", err))
	abortRequest(c, "server error")
}
