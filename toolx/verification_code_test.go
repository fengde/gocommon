package toolx

import "testing"

func TestNewNumberCode(t *testing.T) {
	t.Log(NewNumberCode(2))
	t.Log(NewNumberCode(4))
	t.Log(NewNumberCode(5))
	t.Log(NewNumberCode(6))
	t.Log(NewNumberCode(30))
}

func TestNewCharCode(t *testing.T) {
	t.Log(NewCharCode(2))
	t.Log(NewCharCode(4))
	t.Log(NewCharCode(8))
	t.Log(NewCharCode(20))
}
