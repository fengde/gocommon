package base64x

import "encoding/base64"

// 通用的base64编码
func Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

// 通用的base64解码
func Decode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

// url encode
func UrlEncode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

// url decode
func UrlDecode(s string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(s)
}
