package encryptx

import (
	"crypto/md5"
	"fmt"
)

// Md5 返回md5字节序列
func Md5(data []byte) []byte {
	digest := md5.New()
	digest.Write(data)
	return digest.Sum(nil)
}

// Md5Hex 返回md5 16进制序列
func Md5Hex(data []byte) string {
	return fmt.Sprintf("%x", Md5(data))
}
