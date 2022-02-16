package cachex

import (
	"golang.org/x/sync/singleflight"
)

// CacheFlushHelper 用于防止缓存穿透下，大量重复刷新缓存的请求瞬时打到DB的场景
type CacheFlushHelper struct {
	g  singleflight.Group
}

// NewCacheFlushHelper 创建CacheFlushHelper
func NewCacheFlushHelper() *CacheFlushHelper {
	return &CacheFlushHelper{
		g: singleflight.Group{},
	}
}

// GetFromDB 并发请求到DB,只会实际执行一次fn,然后将结果分享给到其他调用
// key 指定缓存key值，相同key值的并发情况下，只会请求一次fn
// 返回参数：interface, error 同fn返回
func (p *CacheFlushHelper) GetFromDB(key string, fn func() (interface{}, error)) (interface{}, error) {
	inter, err , _ := p.g.Do(key, fn)
	return inter, err
}