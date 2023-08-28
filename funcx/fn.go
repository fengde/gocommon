package funcx

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	goredislib "github.com/redis/go-redis/v9"
)

type Fn struct {
	_fn           func() error
	recoverFlag   bool
	lockFlag      bool
	lockName      string
	lockBlock     bool
	lockTimeout   time.Duration
	redsyncer     *redsync.Redsync
	retryFlag     bool
	retryCount    uint
	retryInterval time.Duration
}

// NewFn 创建函数对象，可支持重试，异常自动捕获，加分布式锁保护
func NewFn(_fn func() error) *Fn {
	return &Fn{
		_fn: _fn,
	}
}

type DoResult struct {
	// 执行耗时
	T time.Duration
	// 是否panic过
	PanicEver bool
	// panic 堆栈信息
	PanicStack string
	// 是否上锁过
	LockEver bool
	// 重试次数
	RetryCount uint
}

// Do 执行函数，带功能参数
func (f *Fn) Do(options ...Option) (dr DoResult, err error) {
	fCopy := Fn{
		_fn: f._fn,
	}

	defer func(start time.Time) {
		dr.T = time.Since(start)
	}(time.Now())

	for _, option := range options {
		option(&fCopy)
	}

	if fCopy.lockFlag {
		mutex := fCopy.redsyncer.NewMutex(fCopy.lockName, redsync.WithExpiry(fCopy.lockTimeout))

		for {
			err = mutex.Lock()
			if err != nil {
				if fCopy.lockBlock {
					time.Sleep(100 * time.Millisecond)
					continue
				}
				return
			}

			break
		}

		dr.LockEver = true

		unlock, cancel := context.WithCancel(context.Background())
		defer func() {
			mutex.Unlock()
			cancel()
		}()

		go func() {
			ticker := time.NewTicker(fCopy.lockTimeout / 2)
			defer ticker.Stop()

			for {
				select {
				case <-ticker.C:
					ok, err := mutex.Extend()
					if !ok || err != nil {
						return
					}
				case <-unlock.Done():
					return
				}
			}
		}()
	}

	var try uint = 0

	for {
		if fCopy.recoverFlag {
			func() {
				defer func() {
					if e := recover(); e != nil {
						dr.PanicEver = true
						dr.PanicStack = string(debug.Stack())
						err = fmt.Errorf("%v", e)
					}
				}()
				err = fCopy._fn()
			}()
		} else {
			err = fCopy._fn()
		}

		if err != nil {
			if fCopy.retryFlag && try < fCopy.retryCount {
				time.Sleep(fCopy.retryInterval)
				try++
				continue
			}
		}

		break
	}

	dr.RetryCount = try

	return
}

type Option func(f *Fn)

func WithRecover() Option {
	return func(f *Fn) {
		f.recoverFlag = true
	}
}

func WithRetry(count uint, interval time.Duration) Option {
	return func(f *Fn) {
		f.retryFlag = true
		f.retryCount = count
		f.retryInterval = interval
	}
}

func WithSynclock(name string, redisCli *goredislib.Options, block bool) Option {
	return func(f *Fn) {
		f.lockFlag = true
		f.lockName = name
		f.lockBlock = block
		f.lockTimeout = 5 * time.Second
		f.redsyncer = redsync.New(goredis.NewPool(goredislib.NewClient(redisCli)))
	}
}
