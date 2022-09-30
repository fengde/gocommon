package taskx

import (
	"fmt"
)

// 串行执行器，逐个执行函数，直到报错为止退出
type SerialTaskGroup struct {
	funcs []func() error
}

func NewSerialTaskGroup(funcs ...func() error) *SerialTaskGroup {
	return &SerialTaskGroup{
		funcs: funcs,
	}
}

func (p *SerialTaskGroup) Run() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recover: %v", r)
		}
	}()

	for _, f := range p.funcs {
		if err = f(); err != nil {
			return
		}
	}

	return nil
}
