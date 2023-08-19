package server

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/adharshmk96/stk-template/singlemod/internals/http/handler"
	"github.com/adharshmk96/stk-template/singlemod/internals/service"
	"github.com/adharshmk96/stk-template/singlemod/internals/storage/sqlite"
	"github.com/adharshmk96/stk-template/singlemod/server/infra"
	svrmw "github.com/adharshmk96/stk-template/singlemod/server/middleware"
	"github.com/adharshmk96/stk-template/singlemod/server/routing"
	"github.com/adharshmk96/stk/gsk"
	"github.com/adharshmk96/stk/pkg/db"
	"github.com/adharshmk96/stk/pkg/middleware"
)

func StartHttpServer(port string) (*gsk.Server, chan bool) {

	logger := infra.GetLogger()

	serverConfig := &gsk.ServerConfig{
		Port:   port,
		Logger: logger,
	}

	server := gsk.New(serverConfig)

	rateLimiter := svrmw.RateLimiter()
	server.Use(rateLimiter)
	server.Use(middleware.RequestLogger)
	server.Use(middleware.CORS(middleware.CORSConfig{
		AllowAll: true,
	}))

	infra.LoadDefaultConfig()

	intializeServer(server)

	server.Start()

	// graceful shutdown
	done := make(chan bool)

	// A go routine that listens for os signals
	// it will block until it receives a signal
	// once it receives a signal, it will shutdown close the done channel
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		if err := server.Shutdown(); err != nil {
			logger.Error(err.Error())
		}

		close(done)
	}()

	return server, done
}

func intializeServer(server *gsk.Server) {
	conn := db.GetSqliteConnection("sqlite.db")

	stktemplateStorage := sqlite.NewSqliteRepo(conn)
	stktemplateService := service.NewPingService(stktemplateStorage)
	stktemplateHandler := handler.NewPingHandler(stktemplateService)

	routing.SetupPingRoutes(server, stktemplateHandler)
}
