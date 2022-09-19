package ratelimitx

// 漏桶算法限速器
type Ratelimit LeakyBucketRatelimit

// 创建漏桶算法限速器，指定速率(每秒）
func NewRatelimit(rps int) *Ratelimit {
	return (*Ratelimit)(NewLeakyBucketRatelimit(rps))
}

// 运行函数，如果并发执行，将严格执行限速器
func (p *Ratelimit) Run(f func()) {
	p.limiter.Take()
	f()
}
