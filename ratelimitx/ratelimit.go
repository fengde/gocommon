package ratelimitx

import (
	"time"

	"go.uber.org/ratelimit"
)

type Ratelimit struct {
	limiter ratelimit.Limiter
}

// NewRatelimit 创建限速器，指定速率(每秒）
func NewRatelimit(rps int) *Ratelimit {
	return &Ratelimit{
		limiter: ratelimit.New(rps),
	}
}

// NewRatelimitCommon 创建通用的限速器，指定per时间内可以处理的次数
func NewRatelimitCommon(r int, per time.Duration) *Ratelimit {
	return &Ratelimit{
		limiter: ratelimit.New(r, ratelimit.Per(per)),
	}
}

// Run 运行函数，如果并发执行，将严格执行限速器
func (p *Ratelimit) Run(f func()) {
	p.limiter.Take()
	f()
}
