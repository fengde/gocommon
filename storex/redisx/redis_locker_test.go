package redisx

import (
	"context"
	"testing"
)

func TestLocker_Lock(t *testing.T) {
	locker := NewLocker(client, "test1", 10)
	ok, err := locker.Lock(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	if ok {
		t.Log("上锁成功")
	} else {
		t.Log("上锁失败")
	}
}

func TestLocker_Unlock(t *testing.T) {
	locker := NewLocker(client, "test2", 10)
	ok, err := locker.Lock(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	if ok {
		t.Log("上锁成功")
	} else {
		t.Log("上锁失败")
	}
	if err := locker.Unlock(context.Background()); err != nil {
		t.Error(err)
		return
	}
}

func TestNewLocker(t *testing.T) {
	NewLocker(client, "fengde", 10)
}
