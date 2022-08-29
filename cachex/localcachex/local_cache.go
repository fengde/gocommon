package localcachex

// 本地localcachex是对 https://github.com/coocood/freecache 的二次封装

import (
	"runtime/debug"

	"github.com/coocood/freecache"
)

type LocalCache struct {
	cacheSize int
	cache     *freecache.Cache
}

// NewLocalCache 新建LocalCache, 默认缓存大小512KB
func NewLocalCache() *LocalCache {
	return NewLocalCacheWithSize(1024 * 512)
}

// NewLocalCacheWithSize 指定缓存大小来新建LocalCache
// 参数:
// 	cacheSize: 指定本地缓存的大小，单位byte，最小512KB, 如果小于512KB将强制成512KB
//  gcPercent: 如果设置的缓存比较大，可以设置合理的gc回收频率
func NewLocalCacheWithSize(cacheSize int, gcPercent ...int) *LocalCache {
	if len(gcPercent) > 0 && gcPercent[0] > 0 && gcPercent[0] < 100 {
		debug.SetGCPercent(gcPercent[0])
	}
	if cacheSize < 1024*512 {
		cacheSize = 1024 * 512
	}
	return &LocalCache{
		cacheSize: cacheSize,
		cache:     freecache.NewCache(cacheSize),
	}
}

// Set 插入kv，key不过期
// 注意：
// 	如果缓存满了，历史key可能被新key取代
func (p *LocalCache) Set(key []byte, value []byte) error {
	return p.SetWithExpire(key, value, 0)
}

// SetWithExpire 插入kv, 设置过期时间
// 参数：
// 	expireSeconds 如果 <=0，表示key不过期
// 注意：
// 	如果缓存满了，历史key可能被新key取代
func (p *LocalCache) SetWithExpire(key []byte, value []byte, expireSeconds int) error {
	return p.cache.Set(key, value, expireSeconds)
}

// Get 查询key
// 返回参数：
// 	err: 如果未查找到key, err不为nil
func (p *LocalCache) Get(key []byte) (value []byte, err error) {
	return p.cache.Get(key)
}

// Del 删除key
// 返回参数：
// 	bool: 命中删除，bool->true; 否则false
func (p *LocalCache) Del(key []byte) bool {
	return p.cache.Del(key)
}

// ClearAll 清空缓存
func (p *LocalCache) ClearAll() {
	p.cache.Clear()
}

// HitRate 缓存命中率
func (p *LocalCache) HitRate() float64 {
	return p.cache.HitRate()
}

// TTL key剩下的TTL时长
func (p *LocalCache) TTL(key []byte) (timeLeft uint32, err error) {
	return p.cache.TTL(key)
}
