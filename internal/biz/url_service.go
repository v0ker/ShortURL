package biz

import (
	"ShortURL/internal/api/handler"
	"ShortURL/internal/config"
	"ShortURL/internal/types"
	"ShortURL/internal/utils"
	"context"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type UrlData interface {
	Create(ctx context.Context, url *types.UrlRecord) error
	GetByCode(ctx context.Context, code int64) (*types.UrlRecord, error)
}

type UrlService struct {
	urlData   UrlData
	idData    IdData
	log       *zap.Logger
	urlConfig config.UrlConfig
}

func NewUrlService(data UrlData, idData IdData, config *config.Configuration, log *zap.Logger) handler.UrlService {
	return &UrlService{
		urlData:   data,
		idData:    idData,
		urlConfig: config.Url,
		log:       log,
	}
}

func (u UrlService) ShortenUrl(ctx context.Context, url string, ttl int32) (string, error) {
	// query recent url from cache (a short time window)
	// if exists, return same short url
	// else, create new short url
	codeId, err := u.idData.GetId()
	if err != nil {
		return "", err
	}
	var urlRecord = &types.UrlRecord{
		Url:     url,
		Code:    codeId,
		Ttl:     ttl,
		Created: time.Now(),
	}
	err = u.urlData.Create(ctx, urlRecord)
	shortenUrl := fmt.Sprintf("%s/s/%s", u.urlConfig.Domain, utils.Int2String(codeId, u.urlConfig.MinLength))
	return shortenUrl, err
}

func (u UrlService) ExpandUrl(ctx context.Context, url string) (*types.UrlRecord, error) {
	//TODO implement me
	panic("implement me")
}
