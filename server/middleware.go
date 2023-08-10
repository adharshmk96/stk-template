package server

import (
	"time"

	"github.com/adharshmk96/stk/gsk"
	"github.com/adharshmk96/stk/pkg/middleware"
)

func rateLimiter() gsk.Middleware {
	rlConfig := middleware.RateLimiterConfig{
		RequestsPerInterval: 10,
		Interval:            60 * time.Second,
	}
	rateLimiter := middleware.NewRateLimiter(rlConfig)
	return rateLimiter.Middleware
}
