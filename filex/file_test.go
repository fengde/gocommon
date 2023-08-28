package filex

import (
	"strings"
	"testing"
)

func TestWriteAppend(t *testing.T) {
	WriteAppend("./write_test.log", strings.NewReader("hello"))
}
