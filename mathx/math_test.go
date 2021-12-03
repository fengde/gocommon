package mathx

import "testing"

func TestRand(t *testing.T) {
	t.Log(Rand(-1, 0))
	t.Log(Rand(-1, 10))
	t.Log(Rand(-8, -1))
	t.Log(Rand(1, 5))
	t.Log(Rand(1, 1))
}
