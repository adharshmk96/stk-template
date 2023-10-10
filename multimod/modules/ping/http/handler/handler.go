package handler

import (
	"github.com/adharshmk96/stk-template/multimod/internals/ping/interfaces"
)

type pingHandler struct {
	pingService interfaces.PingService
}

func NewPingHandler(pingService interfaces.PingService) interfaces.PingHandlers {
	return &pingHandler{
		pingService: pingService,
	}
}
