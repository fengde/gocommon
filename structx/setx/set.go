package setx

import "sync"

// Set 简单的集合功能实现
type Set[T comparable] struct {
	m sync.Map
}

// NewSet 新建set对象
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{}
}

// Has 判断集合中是否存在值，返回bool
func (p *Set[T]) Has(v T) bool {
	_, ok := p.m.Load(v)
	return ok
}

// NotHas 判断集合中是否不存在值，返回bool
func (p *Set[T]) NotHas(v T) bool {
	return !p.Has(v)
}

// Store 存值，args支持传入多个值
func (p *Set[T]) Store(args ...T) {
	for _, v := range args {
		p.m.Store(v, 1)
	}
}

// Delete 删除值，args支持删除多个值
func (p *Set[T]) Delete(args ...T) {
	for _, v := range args {
		p.m.Delete(v)
	}
}

// Items 获取Set所有值，返回无序数组
func (p *Set[T]) Items() []T {
	var items []T
	p.m.Range(func(key, _ any) bool {
		items = append(items, key.(T))
		return true
	})
	return items
}

// LeftDifference 想象A，B两个圈，公共区域相交在一起，此函数返回A圈特有的部分
func (p *Set[T]) LeftDifference(b *Set[T]) []T {
	var left []T
	for _, v := range p.Items() {
		if !b.Has(v) {
			left = append(left, v)
		}
	}
	return left
}

// RightDifference 想象A，B两个圈，公共区域相交在一起，此函数返回B圈特有的部分
func (p *Set[T]) RightDifference(b *Set[T]) []T {
	var right []T
	for _, v := range b.Items() {
		if !p.Has(v) {
			right = append(right, v)
		}
	}
	return right
}

// InnerHave 想象A，B两个圈，公共区域相交在一起，此函数返回A、B共有的部分
func (p *Set[T]) InnerHave(b *Set[T]) []T {
	var inner []T
	for _, v := range p.Items() {
		if b.Has(v) {
			inner = append(inner, v)
		}
	}
	return inner
}
