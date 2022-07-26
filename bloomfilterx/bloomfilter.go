package bloomfilterx

import (
	bloom "github.com/bits-and-blooms/bloom/v3"
)

type BloomFilter struct {
	bf *bloom.BloomFilter
}

// NewBloomFilter 创建布隆过滤器，指定大概要装入的key数，以及能够接受的错误率，比如0.01错误率（即100个里面可能出现一个误差）
func NewBloomFilter(n uint, fp float64) *BloomFilter {
	return &BloomFilter{
		bf: bloom.NewWithEstimates(n, fp),
	}
}

// Add 新增key
func (p *BloomFilter) Add(key string) {
	p.bf.AddString(key)
}

// IsExist 判断key是否存在，需要注意的是, 由于布隆过滤器本身特点，检测返回true不能百分百代表key存在（有可能发生了碰撞），返回false可以百分百代表key不存在
func (p *BloomFilter) IsExist(key string) bool {
	return p.bf.TestString(key)
}
