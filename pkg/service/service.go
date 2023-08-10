package service

import (
	"github.com/adharshmk96/stk-template/pkg/core"
)

type pingService struct {
	pingStorage core.PingStorage
}

func NewPingService(storage core.PingStorage) core.PingService {
	return &pingService{
		pingStorage: storage,
	}
}
