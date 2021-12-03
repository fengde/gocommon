package sysx

import (
	"os"

	"github.com/fengde/gocommon/logx"
)

// Hostname 返回主机名
func Hostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		logx.Warn(err)
		hostname = ""
	}
	return hostname
}
