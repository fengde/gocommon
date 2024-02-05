package safex

import (
	"testing"
)

func TestRecover(t *testing.T) {
	defer Recover()
	panic("aaaaa")
}
