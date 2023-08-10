package server

import (
	"github.com/adharshmk96/stk-template/pkg/http/handler"
	"github.com/adharshmk96/stk-template/pkg/service"
	"github.com/adharshmk96/stk-template/pkg/storage/sqlite"
	"github.com/adharshmk96/stk/gsk"
	"github.com/adharshmk96/stk/pkg/db"
)

func setupRoutes(server *gsk.Server) {

	conn := db.GetSqliteConnection("sqlite.db")

	stktemplateStorage := sqlite.NewSqliteRepo(conn)
	stktemplateService := service.NewPingService(stktemplateStorage)
	stktemplateHandler := handler.NewPingHandler(stktemplateService)

	server.Get("/ping", stktemplateHandler.PingHandler)
}
