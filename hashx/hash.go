package hashx

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"

	"github.com/spaolacci/murmur3"
)

// Hash 返回hash值
func Hash(data []byte) uint64 {
	return murmur3.Sum64(data)
}

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
