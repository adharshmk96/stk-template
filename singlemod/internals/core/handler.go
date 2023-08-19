package core

import "github.com/adharshmk96/stk/gsk"

type PingHandlers interface {
	PingHandler(gc *gsk.Context)
}
