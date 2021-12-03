package mathx

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// MaxInt64 返回较大的数
func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}

	return b
}

// MinInt64 返回较小的数
func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}

	return b
}

// Rand 返回[min, max]之间的随机数n ( min<=n<=max)，min支持负数
func Rand(min, max int64) int64 {
	if min > max {
		panic("Rand min>max!")
	}
	return rand.Int63n(max+1-min) + min
}
