package service

import (
	"github.com/adharshmk96/stk-template/singlemod/internals/core"
)

type pingService struct {
	pingStorage core.PingStorage
}

func NewPingService(storage core.PingStorage) core.PingService {
	return &pingService{
		pingStorage: storage,
	}
}
