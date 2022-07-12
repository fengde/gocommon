package toolx

import (
	"fmt"

	"github.com/fengde/gocommon/mathx"
)

// NewNumberCode 创建指定长度的数字验证码
func NewNumberCode(length int) string {
	var s string
	for i := 0; i < length; i++ {
		s += fmt.Sprintf("%d", mathx.Rand(0, 9))
	}

	return s
}

// NewCharCode 创建指定长度的字符验证码
func NewCharCode(length int) string {
	var code string
	for i := 0; i < length; i++ {
		n := mathx.Rand(97, 122)
		code += string(byte(n))
	}
	return code
}
