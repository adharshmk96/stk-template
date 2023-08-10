package core

import "github.com/adharshmk96/stk/gsk"

type PingHandler interface {
	PingHandler(gc *gsk.Context)
}
