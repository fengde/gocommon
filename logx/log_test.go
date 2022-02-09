package logx

import (
	"testing"
)

func TestDebug(t *testing.T) {
	ctx := NewCtx()
	DebugWithCtx(ctx, "hello world")
}

func TestSetLogFile(t *testing.T) {
	SetLogFile("./test.log", 2)
	DebugWithCtx(nil, "hello world")
}

func TestSentryHook(t *testing.T) {
	AddSentryHook("your dsn", []Level{
		ErrorLevel,
	})

	Error("test sentry")
	ErrorfWithCtx(NewCtx(), "test sentry2")
}