package ratelimitx

import (
	"testing"
	"time"

	"github.com/fengde/gocommon/timex"
)

func TestLeakyBucketRatelimit_Run(t *testing.T) {
	rl := NewLeakyBucketRatelimit(3)
	for {
		rl.Run(func() {
			t.Log(timex.NowTimeString(), "hello world")
		})
	}
}

func TestLeakyBucketRatelimit2_Run(t *testing.T) {
	rl := NewLeakyBucketRatelimitCommon(3, 2*time.Second)
	for {
		rl.Run(func() {
			t.Log(timex.NowTimeString(), "hello world")
		})
	}
}
