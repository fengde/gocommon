package cachex

import (
	"github.com/fengde/gocommon/logx"
	"github.com/fengde/gocommon/safex"
	"testing"
	"time"
)

func TestCacheFlushHelper_GetFromDB(t *testing.T) {
	helper := NewCacheFlushHelper()

	fn := func() (interface{}, error){
		logx.Info("查询DB，project_id=4的项目信息")
		return 1, nil
	}

	var i int64

	for i < 100 {
		safex.Go(func() {
			logx.Info(helper.GetFromDB("project_id_4", fn))
		})
		i++
	}

	time.Sleep(time.Second * 100)
}
