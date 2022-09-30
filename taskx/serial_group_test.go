package taskx

import (
	"testing"
)

func TestNewSerialTaskGroup(t *testing.T) {
	var a = func() error {
		t.Log("a")
		return nil
	}
	var b = func() error {
		t.Log("b")
		return nil
	}
	var c = func() error {
		t.Log("c")
		return nil
	}
	var panicHappen = func() error {
		panic("wrong")
	}
	stg := NewSerialTaskGroup(a, b, c, panicHappen)
	if err := stg.Run(); err != nil {
		t.Log(err)
		return
	}
}
