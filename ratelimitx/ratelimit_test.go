package ratelimitx

import (
	"testing"
	"time"
)

func TestRatelimit_Run(t *testing.T) {
	rl := NewRatelimit(2)
	for {
		rl.Run(func() {
			t.Log("hello world")
		})
		time.Sleep(time.Millisecond)
	}
}
