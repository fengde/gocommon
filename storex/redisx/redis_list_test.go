package redisx

import (
	"testing"
)

func TestNewList(t *testing.T) {
	client := NewClient("127.0.0.1:6379", 0, "")
	list := client.NewList("test")
	if err := list.Push("abc"); err != nil {
		t.Error(err)
		return
	}
	value, err := list.Pop()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(value)
}
