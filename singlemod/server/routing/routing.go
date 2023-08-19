package routing

import (
	"github.com/adharshmk96/stk-template/singlemod/internals/core"
	"github.com/adharshmk96/stk/gsk"
)

func SetupPingRoutes(server *gsk.Server, stktemplateHandler core.PingHandlers) {
	server.Get("/ping", stktemplateHandler.PingHandler)
}