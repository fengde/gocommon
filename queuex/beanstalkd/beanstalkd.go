package beanstalkd

import (
	"time"

	gobeanstalk "github.com/beanstalkd/go-beanstalk"
	"github.com/fengde/gocommon/errorx"
	"github.com/fengde/gocommon/safex"
	"github.com/fengde/gocommon/taskx"
)

type Queue struct {
	queue       *gobeanstalk.Tube
	jobHandle   func(id uint64, body []byte) (retry bool)
	ttr         time.Duration
	concurrency int
	close       chan int
}

/* NewQueue 新建队列
参数：
	address - beanstalkd地址
	topic - 队列名称，cube名称
	ttr - callback执行超时，超时任务将回到队列，重新被消费
	concurrency - 并发消费的任务数
	callback - 回调处理函数； unretry控制是否需要重试（将job重新丢回队列）
*/
func NewQueue(address string, topic string, ttr time.Duration, concurrency int, jobHandle func(id uint64, body []byte) (retry bool)) (*Queue, error) {
	c, err := gobeanstalk.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	if concurrency < 1 {
		return nil, errorx.New("concurrency is valild")
	}

	p := &Queue{
		queue:       gobeanstalk.NewTube(c, topic),
		jobHandle:   jobHandle,
		ttr:         ttr,
		concurrency: concurrency,
		close:       make(chan int),
	}

	safex.Go(p.consumer)

	return p, nil
}

// consumer 异步触发消费
func (p *Queue) consumer() {
	g := taskx.NewTaskGroup(int64(p.concurrency))
	tubes := gobeanstalk.NewTubeSet(p.queue.Conn, p.queue.Name)
	for {
		select {
		case <-p.close:
			g.Wait()
			return
		default:
			id, body, err := tubes.Reserve(time.Minute)
			if err == nil {
				g.Run(func() {
					if retry := p.jobHandle(id, body); !retry {
						p.queue.Conn.Delete(id)
					}
				})
			}
		}
	}
}

// SendJob 发送即时任务
func (p *Queue) SendJob(body []byte) (id uint64, err error) {
	return p.queue.Put(body, 1, 0, p.ttr)
}

// SendDelayJob 发送延时任务
func (p *Queue) SendDelayJob(body []byte, delay time.Duration) (id uint64, err error) {
	return p.queue.Put(body, 1, delay, p.ttr)

}

// QueueStats 查看队列状态
func (p *Queue) QueueStats() (map[string]string, error) {
	return p.queue.Stats()
}

// Pause 队列静默
func (p *Queue) Pause(d time.Duration) error {
	return p.queue.Pause(d)
}

// JobStats job状态
func (p *Queue) JobStats(id uint64) (map[string]string, error) {
	return p.queue.Conn.StatsJob(id)
}

// Touch job执行ttr续期
func (p *Queue) Touch(id uint64) error {
	return p.queue.Conn.Touch(id)
}

// Close 关闭队列
func (p *Queue) Close() {
	close(p.close)
	p.queue.Conn.Close()
}
