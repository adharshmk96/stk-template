package routing

import (
	"github.com/adharshmk96/stk-template/multimod/internals/ping/interfaces"
	"github.com/adharshmk96/stk/gsk"
)

func SetupPingRoutes(server *gsk.Server, stktemplateHandler interfaces.PingHandlers) {
	server.Get("/ping", stktemplateHandler.PingHandler)
}
