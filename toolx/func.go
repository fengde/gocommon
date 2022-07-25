package toolx

import (
	"time"
)

// RepeatFunc 重复执行函数f，n次
func RepeatFunc(n int, f func()) {
	for i := 0; i < n; i++ {
		f()
	}
}

// TickerFunc 定时执行函数f, 间隔d时长
func TickerFunc(d time.Duration, f func()) {
	ticker := time.NewTicker(d)

	for range ticker.C {
		f()
	}
}
