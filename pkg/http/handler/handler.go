package handler

import (
	"github.com/adharshmk96/stk-template/pkg/core"
)

type pingHandler struct {
	pingService core.PingService
}

func NewPingHandler(pingService core.PingService) core.PingHandler {
	return &pingHandler{
		pingService: pingService,
	}
}
