package encryptx

import (
	"crypto/sha256"
	"fmt"
)

// Sha256 返回sha256字节序列
func Sha256(data []byte) []byte {
	digest := sha256.New()
	digest.Write(data)
	return digest.Sum(nil)
}

// Sha256Hex 返回sha256 16进制序列
func Sha256Hex(data []byte) string {
	return fmt.Sprintf("%x", Sha256(data))
}
