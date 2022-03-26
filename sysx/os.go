package sysx

import (
	"github.com/fengde/gocommon/logx"
	"os"
	"os/signal"
	"runtime/debug"
	"sync"
	"syscall"
)


var (
	once  sync.Once
	onSig = make([]func(os.Signal), 0, 16)
)

// 当收到操作系统的退出信号时触发
func OnSignalExit(f func(sig os.Signal)) {
	if f != nil {
		once.Do(func() {
			sigs := make(chan os.Signal, 1)
			signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGILL, syscall.SIGABRT, syscall.SIGKILL, syscall.SIGTERM)
			go func() {
				sig := <-sigs
				logx.Infof("recv signal: %v", sig.String())
				for i := len(onSig) - 1; i >= 0; i-- {
					defer func() {
						if e := recover(); e != nil {
							debug.PrintStack()
						}
					}()
					onSig[i](sig)
				}
				os.Exit(int(sig.(syscall.Signal)))
			}()
		})
		onSig = append(onSig, f)
	}
}

