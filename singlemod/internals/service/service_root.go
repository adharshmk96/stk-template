package service

import (
	"github.com/adharshmk96/stk-template/singlemod/internals/core/entity"
)

type pingService struct {
	pingStorage entity.PingStorage
}

func NewPingService(storage entity.PingStorage) entity.PingService {
	return &pingService{
		pingStorage: storage,
	}
}
