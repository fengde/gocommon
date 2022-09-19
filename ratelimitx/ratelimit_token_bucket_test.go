package ratelimitx

import (
	"testing"
	"time"

	"github.com/fengde/gocommon/logx"
)

func TestNewTokenBucketRatelimit_Run(t *testing.T) {
	limiter := NewTokenBucketRatelimit(1, 10)
	for {
		limiter.Run(func() {
			logx.Info("hello world")
		})
	}
}

func TestNewTokenBucketRatelimit2_Run(t *testing.T) {
	limiter := NewTokenBucketRatelimit(1, 10)
	var i = 0
	for i < 100 {
		if err := limiter.RunWithTimeout(time.Millisecond*500, func() {
			logx.Info("hello world")
		}); err != nil {
			logx.Error(err)
		}
		i++
	}
}
