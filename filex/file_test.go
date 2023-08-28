package filex

import (
	"strings"
	"testing"
)

func TestWriteAppend(t *testing.T) {
	WriteAppend("./write_test.log", strings.NewReader("hello"))
}

func TestIsFileExist(t *testing.T) {
	t.Log(IsFileExist("./file_test.go2"))
	t.Log(IsFileExist("./file_test.go"))
	t.Log(IsFileExist("../filex"))
}

func TestSha(t *testing.T) {
	t.Log(Sha("./file.go", 256))
}
