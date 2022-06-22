package beanstalkd

import (
	"testing"
	"time"

	"github.com/fengde/gocommon/logx"
)

func getQueue() *Queue {
	queue, err := NewQueue("127.0.0.1:11300", "test", time.Minute, 2, func(id uint64, body []byte) (retry bool) {
		logx.Info(string(body))
		return
	})
	if err != nil {
		panic(err)
	}

	return queue
}

func TestNewQueue(t *testing.T) {
	queue := getQueue()
	defer queue.Close()

	id, err := queue.SendJob([]byte("hello world"))
	logx.Info(id, err)

	time.Sleep(time.Second * 3)
}

func TestQueue_SendDelayJob(t *testing.T) {
	queue := getQueue()
	defer queue.Close()

	id, err := queue.SendDelayJob([]byte("hello world"), time.Second*3)
	logx.Info(id, err)

	time.Sleep(time.Second * 5)
}
