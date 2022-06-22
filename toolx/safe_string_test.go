package toolx

import (
	"testing"

	"github.com/fengde/gocommon/logx"
)

func TestSafeString(t *testing.T) {
	logx.Info(SafeString("abc"))
	logx.Info(SafeString("abc", 0))
	logx.Info(SafeString("abc", 2))
	logx.Info(SafeString("abc", 3))
	logx.Info(SafeString("abc", 10))
}
