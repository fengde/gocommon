package syncx

import (
	"github.com/fengde/gocommon/logx"
	"testing"
	"time"
)

func TestFuncLocker_Exec(t *testing.T) {
	locker := NewFuncLocker()
	i := 0
	for i < 10 {
		go locker.Exec(func() {
			logx.Info("here")
			time.Sleep(time.Second * 3)
		})
		time.Sleep(time.Millisecond)
		i++
	}
	time.Sleep(time.Minute)
}
