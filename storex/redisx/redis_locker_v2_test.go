package redisx

import (
	"context"
	"testing"
	"time"

	"github.com/fengde/gocommon/logx"
	"github.com/fengde/gocommon/taskx"
)

func TestNewLockerV2(t *testing.T) {
	locker := NewLockerV2(client)

	g := taskx.NewTaskGroup(2)

	index := 0
	for index < 2 {
		g.Run(func() {
			ok, err := locker.LockBlock("fedelfeng", 15, func() {
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()

				for {
					select {
					case <-ctx.Done():
						return
					default:
						logx.Info("ok")
					}
					time.Sleep(time.Second)
				}
			})
			logx.Info(ok, err)
		})
		index++
	}

	g.Wait()
}
