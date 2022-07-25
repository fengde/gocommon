package toolx

import (
	"testing"
	"time"
)

func TestRepeatFunc(t *testing.T) {
	RepeatFunc(2, func() {
		t.Log("hello world")
	})
}

func TestTickerFunc(t *testing.T) {
	TickerFunc(2*time.Second, func() {
		t.Log("hello world")
	})
}
