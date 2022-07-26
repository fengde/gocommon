package bloomfilterx

import (
	"testing"
)

func TestBloomFilter_IsExist(t *testing.T) {
	bf := NewBloomFilter(10*10, 0.00000001)
	bf.Add("hello")
	t.Log(bf.IsExist("hello"))
}
