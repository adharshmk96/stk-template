package handler

import (
	"github.com/adharshmk96/stk-template/singlemod/internals/core/entity"
)

type pingHandler struct {
	service entity.PingService
}

func NewPingHandler(service entity.PingService) entity.PingHandlers {
	return &pingHandler{
		service: service,
	}
}
