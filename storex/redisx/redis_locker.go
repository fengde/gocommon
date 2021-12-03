package redisx

import (
	"fmt"
	"time"

	"github.com/fengde/gocommon/errorx"
	"github.com/fengde/gocommon/verificationCodex"
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
		value:            fmt.Sprintf("%v.%v", time.Now().UnixNano(), verificationCodex.NewNumberCode(4)),
		autoUnlockSecond: autoUnlockSecond,
	}
}

// Lock 上锁，非阻塞，如果上锁成功，bool=true
func (p *Locker) Lock() (bool, error) {
	err := p.client.client.Do(p.client.getCtx(), "SET", p.key, p.value, "EX", p.autoUnlockSecond, "NX").Err()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, errorx.WithStack(err)
	}
	return true, nil
}

// Unlock 释放锁
func (p *Locker) Unlock() error {
	value, err := p.client.GetString(p.key)
	if err != nil {
		return errorx.WithStack(err)
	}
	if value == p.value {
		if err := p.client.Del(p.key); err != nil {
			return errorx.WithStack(err)
		}
	}

	return nil
}
