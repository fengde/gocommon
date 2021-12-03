package verificationCodex

import (
	"fmt"
	"math"

	"github.com/fengde/gocommon/mathx"
)

// NewNumberCode 创建指定长度的数字验证码
func NewNumberCode(length int) string {
	min := int64(math.Pow10(length - 1))
	max := int64(math.Pow10(length))
	return fmt.Sprintf("%d", mathx.Rand(min, max-1))
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
