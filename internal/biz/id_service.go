package biz

import (
	"ShortURL/internal/api/handler"
	"go.uber.org/zap"
)

type IdData interface {
	GetId() (int64, error)
}

type IdService struct {
	log    *zap.Logger
	idData IdData
}

func NewIdService(log *zap.Logger, idData IdData) handler.IdService {
	return &IdService{
		log:    log,
		idData: idData,
	}
}

func (i IdService) NextId() (int64, error) {
	return i.idData.GetId()
}
