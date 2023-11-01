package data

import (
	"ShortURL/internal/types"
	"ShortURL/internal/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UrlRecordData struct {
	data *Data
	log  *zap.Logger
}

func NewUrlRecordData(data *Data, log *zap.Logger) *UrlRecordData {
	return &UrlRecordData{
		data: data,
		log:  log,
	}
}

func (u UrlRecordData) Create(urlRecord *types.UrlRecord) error {
	result := u.data.db.Create(urlRecord)
	return result.Error
}

func (u UrlRecordData) GetByCode(ctx context.Context, code int64) (*types.UrlRecord, error) {
	urlRecord, err := u.getByCodeFomCache(ctx, code)
	if err != nil {
		return nil, err
	}
	value := u.data.db.Where("`code` = ?", code).First(&urlRecord)
	if value.Error != nil && value.Error != gorm.ErrRecordNotFound {
		return nil, value.Error
	}
	if value.Error == gorm.ErrRecordNotFound {
		urlRecord = nil
	}
	_ = u.setUrlRecordToCache(ctx, urlRecord)
	return urlRecord, nil
}

func (u UrlRecordData) getByCodeFomCache(ctx context.Context, code int64) (*types.UrlRecord, error) {
	key := u.codeCacheKey(code)
	cmd := u.data.rdb.Get(ctx, key)
	if cmd.Err() == nil || cmd.Err() == redis.Nil {
		if cmd.Err() == redis.Nil {
			return nil, nil
		}
		if cmd.Val() == utils.EmptyCache {
			return nil, nil
		}
		urlRecord, err := u.unmarshalUrlRecord(cmd.Val())
		return urlRecord, err
	} else {
		return nil, cmd.Err()
	}
}

func (u UrlRecordData) setUrlRecordToCache(ctx context.Context, urlRecord *types.UrlRecord) error {
	if urlRecord == nil {
		return u.data.rdb.Set(ctx, u.codeCacheKey(urlRecord.Code), utils.EmptyCache, utils.CacheDuration).Err()
	}
	data, _ := json.Marshal(urlRecord)
	return u.data.rdb.Set(ctx, u.codeCacheKey(urlRecord.Code), data, utils.CacheDuration).Err()
}

func (u UrlRecordData) unmarshalUrlRecord(data string) (*types.UrlRecord, error) {
	var urlRecord types.UrlRecord
	err := json.Unmarshal([]byte(data), &urlRecord)
	return &urlRecord, err
}

func (u UrlRecordData) codeCacheKey(code int64) string {
	return fmt.Sprintf("code:%d", code)
}
