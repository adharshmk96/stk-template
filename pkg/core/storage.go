package core

type PingStorage interface {
	Ping() error
}
