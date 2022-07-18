package safex

import "github.com/fengde/gocommon/logx"

// Func 安全执行，内部已处理异常捕获
func Func(fn func()) {
	func() {
		defer Recover()
		fn()
	}()
}

// Go 执行并发协程，内部已处理异常捕获
func Go(fn func()) {
	go Func(fn)
}

// Recover 封装了语言recover，支持传入扫尾函数
func Recover(cleanups ...func()) {
	if p := recover(); p != nil {
		logx.Warn(p)
		for _, cleanup := range cleanups {
			cleanup()
		}
	}
}
