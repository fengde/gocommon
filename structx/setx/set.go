package setx

import "sync"

// Set 简单的集合功能实现
type Set struct {
	m sync.Map
}

// Has 判断集合中是否存在值，返回bool
func (p *Set) Has(v interface{}) bool {
	_, ok := p.m.Load(v)
	return ok
}

// Store 存值，args支持传入多个值
func (p *Set) Store(args ...interface{}) {
	for _, v := range args {
		p.m.Store(v, 1)
	}
}

// Delete 删除值，args支持删除多个值
func (p *Set) Delete(args ...interface{}) {
	for _, v := range args {
		p.m.Delete(v)
	}
}

// Items 获取Set所有值，返回无序数组
func (p *Set) Items() []interface{} {
	var items []interface{}
	p.m.Range(func(key, _ interface{}) bool {
		items = append(items, key)
		return true
	})
	return items
}

// LeftDifference 想象A，B两个圈，公共区域相交在一起，此函数返回A圈特有的部分
func (p *Set) LeftDifference(b Set) []interface{} {
	var left []interface{}
	for _, v := range p.Items() {
		if !b.Has(v) {
			left = append(left, v)
		}
	}
	return left
}

// RightDifference 想象A，B两个圈，公共区域相交在一起，此函数返回B圈特有的部分
func (p *Set) RightDifference(b Set) []interface{} {
	var right []interface{}
	for _, v := range b.Items() {
		if !p.Has(v) {
			right = append(right, v)
		}
	}
	return right
}

// InnerHave 想象A，B两个圈，公共区域相交在一起，此函数返回A、B共有的部分
func (p *Set) InnerHave(b Set) []interface{} {
	var inner []interface{}
	for _, v := range p.Items() {
		if b.Has(v) {
			inner = append(inner, v)
		}
	}
	return inner
}