package redisx

import (
	"github.com/fengde/gocommon/safex"
)

type LockerV2 struct {
	client *Client
}

// NewLockerV2 新建分布式锁
func NewLockerV2(client *Client) *LockerV2 {
	return &LockerV2{
		client: client,
	}
}

// Lock 执行函数，非阻塞
// 返回值：锁成功 返回true，nil; 锁失败 返回false, err
func (p *LockerV2) Lock(sourceID string, autoUnlockSecond int64, fn func()) (bool, error) {
	locker := NewLocker(p.client, sourceID, autoUnlockSecond)
	ok, err := locker.Lock()
	if err != nil || !ok {
		return ok, err
	}

	defer locker.Unlock()

	safex.Func(fn)

	return true, nil
}

// LockBlock 执行函数，阻塞
// 返回值：锁成功 返回true，nil; 锁失败 返回false, err
func (p *LockerV2) LockBlock(sourceID string, autoUnlockSecond int64, fn func()) (bool, error) {
	locker := NewLocker(p.client, sourceID, autoUnlockSecond)

	ok, err := locker.LockBlock()
	if err != nil || !ok {
		return ok, err
	}

	defer locker.Unlock()

	safex.Func(fn)

	return true, nil
}
