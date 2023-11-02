package handler

import (
	"ShortURL/internal/types"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UrlService interface {
	ShortenUrl(ctx context.Context, url string, ttl int32) (string, error)
	ExpandUrl(ctx context.Context, shortenId string) (*types.UrlRecord, error)
}

type UrlHandler struct {
	log        *zap.Logger
	urlService UrlService
}

func NewUrlHandler(log *zap.Logger, service UrlService) *UrlHandler {
	return &UrlHandler{
		log:        log,
		urlService: service,
	}
}

func (u UrlHandler) ShortenUrl(ctx *gin.Context) {
	var urlRequest UrlRequest
	err := ctx.ShouldBindJSON(&urlRequest)
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}
	url, err := u.urlService.ShortenUrl(ctx, urlRequest.Url, urlRequest.Ttl)
	if err != nil {
		ctx.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": url,
	})
}

func (u UrlHandler) ExpandUrl(ctx *gin.Context) {
	id := ctx.Param("id")
	url, err := u.urlService.ExpandUrl(ctx, id)
	if err != nil {
		ctx.HTML(500, "error.html", "Server Error")
	} else {
		ctx.Redirect(302, url.Url)
	}
}

type UrlRequest struct {
	Url string `json:"url" binding:"required"`
	Ttl int32  `json:"ttl" binding:"required"`
}
