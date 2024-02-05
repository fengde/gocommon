package encryptx

import (
	"github.com/spaolacci/murmur3"
)

// Hash 返回hash值
func Hash(data []byte) uint64 {
	return murmur3.Sum64(data)
}
