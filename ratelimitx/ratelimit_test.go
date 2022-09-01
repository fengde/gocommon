package ratelimitx

import (
	"testing"
	"time"

	"github.com/fengde/gocommon/timex"
)

func TestRatelimit_Run(t *testing.T) {
	rl := NewRatelimit(3)
	for {
		rl.Run(func() {
			t.Log(timex.NowTimeString(), "hello world")
		})
	}
}

func TestRatelimit2_Run(t *testing.T) {
	rl := NewRatelimitCommon(3, 2*time.Second)
	for {
		rl.Run(func() {
			t.Log(timex.NowTimeString(), "hello world")
		})
	}
}
