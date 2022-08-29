package funcx

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	Repeat(10, func() {
		t.Log("hello world")
	})
}
