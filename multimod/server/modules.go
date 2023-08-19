package server

import (
	"github.com/adharshmk96/stk-template/multimod/internals/ping/http/handler"
	"github.com/adharshmk96/stk-template/multimod/internals/ping/service"
	"github.com/adharshmk96/stk-template/multimod/internals/ping/storage/sqlite"
	"github.com/adharshmk96/stk-template/multimod/server/routing"
	"github.com/adharshmk96/stk/gsk"
	"github.com/adharshmk96/stk/pkg/db"
)

func intializePing(server *gsk.Server) {
	conn := db.GetSqliteConnection("sqlite.db")

	stktemplateStorage := sqlite.NewSqliteRepo(conn)
	stktemplateService := service.NewPingService(stktemplateStorage)
	stktemplateHandler := handler.NewPingHandler(stktemplateService)

	routing.SetupPingRoutes(server, stktemplateHandler)
}

func initModules(server *gsk.Server) {
	intializePing(server)
}
