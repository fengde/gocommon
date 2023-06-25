package redisx

import (
	"context"
	"testing"
)

func TestNewList(t *testing.T) {
	ctx := context.Background()
	client, _ := NewClient(ctx, "127.0.0.1:6379", 0, "")
	list := client.NewList("test")
	if err := list.Push(ctx, "abc"); err != nil {
		t.Error(err)
		return
	}
	value, err := list.Pop(ctx)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}
