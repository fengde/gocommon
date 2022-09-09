package taskx

import (
	"sync"

	"github.com/fengde/gocommon/safex"
)

// TaskGroup 封装了sync.WaitGroup的任务组
type TaskGroup struct {
	waitGroup sync.WaitGroup
	limitChan chan int64
}

// NewTaskGroup 创建一个任务组，支持传入同时并发的任务数
func NewTaskGroup(concurrency ...int64) *TaskGroup {
	g := TaskGroup{
		limitChan: nil,
	}
	if len(concurrency) > 0 {
		g.limitChan = make(chan int64, concurrency[0])
	}
	return &g
}

// Run 执行任务
func (p *TaskGroup) Run(fn func()) {
	if p.limitChan != nil {
		p.limitChan <- 1
	}

	p.waitGroup.Add(1)

	safex.Go(func() {
		defer func() {
			p.waitGroup.Done()
			if p.limitChan != nil {
				<-p.limitChan
			}
		}()
		fn()
	})
}

// Wait 阻塞等待任务都执行完成
func (p *TaskGroup) Wait() {
	p.waitGroup.Wait()
}
