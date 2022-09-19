package ratelimitx

import (
	"context"
	"time"

	"golang.org/x/time/rate"
)

// 令牌桶算法限速器
type TokenBucketRatelimit struct {
	limiter *rate.Limiter
}

// 创建令牌桶算法限速器，指定每秒投递令牌数r； 令牌桶大小b，初始化令牌桶已装满令牌
func NewTokenBucketRatelimit(r int, b int) *TokenBucketRatelimit {
	return &TokenBucketRatelimit{
		limiter: rate.NewLimiter(rate.Limit(r), b),
	}
}

// 创建更为通用的令牌桶算法限速器，limit可用rate.Every来指定投递速率；令牌桶大小b，初始化令牌桶已装满令牌
func NewTokenBucketRatelimitCommon(limit rate.Limit, b int) *TokenBucketRatelimit {
	return &TokenBucketRatelimit{
		limiter: rate.NewLimiter(limit, b),
	}
}

// 运行函数
func (p *TokenBucketRatelimit) Run(f func()) {
	p.limiter.Wait(context.Background())
	f()
}

// 带超时等待的逻辑，运行函数；如果等待了timeout时长也没有获取到令牌，则返回err
func (p *TokenBucketRatelimit) RunWithTimeout(timeout time.Duration, f func()) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := p.limiter.Wait(ctx); err != nil {
		return err
	}
	f()
	return nil
}
