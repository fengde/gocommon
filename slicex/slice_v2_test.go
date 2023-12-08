package slicex

import (
	"fmt"
	"testing"
)

func TestForEach(t *testing.T) {
	ForEach([]int{1, 2, 3}, func(n int) {
		fmt.Println(n)
	})
}
