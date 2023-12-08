package safex

import (
	"context"
	"runtime/debug"

	"github.com/fengde/gocommon/logx"
)

// FuncCtx 安全执行，内部已处理异常捕获
func FuncCtx(ctx context.Context, fn func()) {
	func() {
		defer RecoverCtx(ctx)
		fn()
	}()
}

// GoCtx 执行并发协程，内部已处理异常捕获
func GoCtx(ctx context.Context, fn func()) {
	go FuncCtx(ctx, fn)
}

// RecoverCtx 封装了语言recover，支持传入扫尾函数
func RecoverCtx(ctx context.Context, cleanups ...func()) {
	if p := recover(); p != nil {
		logx.ErrorWithCtx(ctx, p)
		logx.ErrorWithCtx(ctx, string(debug.Stack()))
	}
	for _, cleanup := range cleanups {
		cleanup()
	}
}
