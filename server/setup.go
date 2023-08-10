package server

import "github.com/adharshmk96/stk/gsk"

func StartServer(port string) *gsk.Server {

	serverConfig := &gsk.ServerConfig{
		Port: port,
	}

	server := gsk.New(serverConfig)

	setupRoutes(server)

	rateLimiter := rateLimiter()
	server.Use(rateLimiter)

	server.Start()

	return server
}
