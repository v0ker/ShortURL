package data

import (
	"ShortURL/internal/api/handler"
	"go.uber.org/zap"
)

type EchoData struct {
	data *Data
	log  *zap.Logger
}

func NewEchoData(data *Data, log *zap.Logger) handler.EchoData {
	return &EchoData{
		data: data,
		log:  log,
	}
}

func (e EchoData) Echo() string {
	return "hello world"
}
