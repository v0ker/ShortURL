package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type IdService interface {
	NextId() (int64, error)
}

type IdHandler struct {
	log       *zap.Logger
	idService IdService
}

func NewIdHandler(log *zap.Logger, service IdService) *IdHandler {
	return &IdHandler{
		log:       log,
		idService: service,
	}
}

func (h *IdHandler) NextId(ctx *gin.Context) {
	id, err := h.idService.NextId()
	if err != nil {
		ctx.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})
	} else {
		ctx.JSON(200, gin.H{
			"code": 200,
			"data": id,
		})
	}

}
