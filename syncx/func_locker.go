package syncx

import "sync"

// FuncLocker 函数锁，确保函数任何时刻只有一个线程在执行它
type FuncLocker struct {
	locker sync.Mutex
}

// NewFuncLocker 新增函数锁
func NewFuncLocker() *FuncLocker {
	return &FuncLocker{
		locker: sync.Mutex{},
	}
}

// Exec 执行函数
func (p *FuncLocker) Exec(fn func()) {
	p.locker.Lock()
	defer p.locker.Unlock()
	fn()
}
