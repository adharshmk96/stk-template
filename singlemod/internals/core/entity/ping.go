package entity

import "github.com/adharshmk96/stk/gsk"

// Domain
type User struct {
	pong string
}

// Storage
type PingStorage interface {
	Ping() error
}

// Service
type PingService interface {
	PingService() string
}

// Handler
type PingHandlers interface {
	PingHandler(gc *gsk.Context)
}
