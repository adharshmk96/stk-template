package service

import (
	"github.com/adharshmk96/stk-template/multimod/internals/ping/interfaces"
)

type pingService struct {
	pingStorage interfaces.PingStorage
}

func NewPingService(storage interfaces.PingStorage) interfaces.PingService {
	return &pingService{
		pingStorage: storage,
	}
}
