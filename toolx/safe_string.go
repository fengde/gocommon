package toolx

import "unicode/utf8"

// SafeString 安全的返回字符串，如果str长度超过safeLength，返回str[0:safeLength]
func SafeString(str string, safeLength ...int) string {
	if len(safeLength) > 0 {
		length := safeLength[0]
		if utf8.RuneCountInString(str) >= length {
			t := []rune(str)
			return string(t[0:length])
		}
	}
	return str
}
