package logx

import (
	"testing"
)

func TestDebug(t *testing.T) {
	ctx := NewCtx()
	DebugWithCtx(ctx, "hello world")
}

func TestSetLogFile(t *testing.T) {
	SetLogFile("./fengde.log", 2)
	DebugWithCtx(nil, "hello world")
}
