package handler

import (
	"github.com/adharshmk96/stk-template/singlemod/internals/core/entity"
)

type pingHandler struct {
	pingService entity.PingService
}

func NewPingHandler(pingService entity.PingService) entity.PingHandlers {
	return &pingHandler{
		pingService: pingService,
	}
}
