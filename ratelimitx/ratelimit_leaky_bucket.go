package ratelimitx

import (
	"time"

	"go.uber.org/ratelimit"
)

// 漏桶算法限速器
type LeakyBucketRatelimit struct {
	limiter ratelimit.Limiter
}

// 创建漏桶算法限速器，指定速率(每秒）
func NewLeakyBucketRatelimit(rps int) *LeakyBucketRatelimit {
	return &LeakyBucketRatelimit{
		limiter: ratelimit.New(rps),
	}
}

// 创建通用的漏桶算法限速器，指定per时间内可以处理的次数
func NewLeakyBucketRatelimitCommon(r int, per time.Duration) *LeakyBucketRatelimit {
	return &LeakyBucketRatelimit{
		limiter: ratelimit.New(r, ratelimit.Per(per)),
	}
}

// 运行函数，如果并发执行，将严格执行限速器
func (p *LeakyBucketRatelimit) Run(f func()) {
	p.limiter.Take()
	f()
}
