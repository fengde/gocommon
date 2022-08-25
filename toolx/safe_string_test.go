package toolx

import (
	"testing"
)

func TestSafeString(t *testing.T) {
	s := "、：我爱我的家人、、、、、："
	t.Log(SafeString(s, 5))
	s = "woaiwojia"
	t.Log(SafeString(s, 2))
}
