package redisx

import (
	"context"
	"fmt"
	"time"

	"github.com/fengde/gocommon/errorx"
	"github.com/fengde/gocommon/toolx"
	"github.com/go-redis/redis/v8"
)

type Locker struct {
	client           *Client
	key              string
	value            string
	autoUnlockSecond int64
}

// NewLocker 新建分布式锁，sourceID 标识资源id, autoUnlockSecond表示资源自动解锁时间
func NewLocker(client *Client, sourceID string, autoUnlockSecond int64) *Locker {
	return &Locker{
		client:           client,
		key:              "redisx-locker:" + sourceID,
		value:            fmt.Sprintf("%v.%v", time.Now().UnixNano(), toolx.NewNumberCode(4)),
		autoUnlockSecond: autoUnlockSecond,
	}
}

// Lock 非阻塞上锁
// 返回值：锁成功 返回true, nil; 锁失败 返回false
func (p *Locker) Lock(ctx context.Context) (bool, error) {
	err := p.client.client.Do(ctx, "SET", p.key, p.value, "EX", p.autoUnlockSecond, "NX").Err()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, errorx.WithStack(err)
	}
	return true, nil
}

// LockBlock 阻塞上锁
// 返回值：锁成功 返回true, nil; 锁失败 返回false
func (p *Locker) LockBlock(ctx context.Context) (bool, error) {
	for {
		ok, err := p.Lock(ctx)
		if err != nil {
			return false, errorx.WithStack(err)
		}

		if ok {
			return true, nil
		}

		time.Sleep(time.Second)
	}
}

// Unlock 释放锁
func (p *Locker) Unlock(ctx context.Context) error {
	value, err := p.client.GetString(ctx, p.key)
	if err != nil {
		return errorx.WithStack(err)
	}
	if value == p.value {
		if err := p.client.Del(ctx, p.key); err != nil {
			return errorx.WithStack(err)
		}
	}

	return nil
}
