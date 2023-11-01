package handler

import (
	"go.uber.org/zap"
)

type UserHandler struct {
	log      *zap.Logger
	echoData EchoData
}

type EchoData interface {
	Echo() string
}

func NewUserHandler(echoData EchoData, log *zap.Logger) *UserHandler {
	return &UserHandler{
		log: log,
	}
}
