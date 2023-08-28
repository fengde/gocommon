package colorx

import "testing"

func TestWithColorPadding(t *testing.T) {
	t.Log(WithColorPadding("hello world", BgGreen))
}
