package ratelimitx

import "go.uber.org/ratelimit"

type Ratelimit struct {
	limiter ratelimit.Limiter
}

// NewRatelimit 创建限速器，指定速率，使用漏桶算法
func NewRatelimit(rps int) *Ratelimit {
	return &Ratelimit{
		limiter: ratelimit.New(rps),
	}
}

// Run 运行函数，如果并发执行，将严格执行限速器
func (p *Ratelimit) Run(f func()) {
	p.limiter.Take()
	f()
}
