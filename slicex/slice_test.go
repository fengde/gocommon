package slicex

import "testing"

func TestContains(t *testing.T) {
	t.Log(IntContains([]int{1, 2, 4, 5}, 2))
}
