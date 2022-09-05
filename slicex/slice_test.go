package slicex

import (
	"testing"
)

func TestContains(t *testing.T) {
	t.Log(IntContains([]int{1, 2, 4, 5}, 2))
}

func TestRemoveRepeat(t *testing.T) {
	t.Log(StrRemoveRepeat([]string{"a", "b", "a", "c", "c", "b", "d"}))
	t.Log(IntRemoveRepeat([]int{1, 2, 3, 4, 5, 6, 4, 3, 2}))
	t.Log(Int64RemoveRepeat([]int64{1, 2, 3, 4, 5, 6, 4, 3, 2}))
}
