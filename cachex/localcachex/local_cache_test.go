package localcachex

import (
	"github.com/fengde/gocommon/logx"
	"strconv"
	"testing"
	"time"
)

func TestNewLocalCache(t *testing.T) {
	var cache = NewLocalCache()
	var index = 0
	for index < 10 {
		if err := cache.Set([]byte(strconv.Itoa(index)), []byte(strconv.Itoa(index))); err != nil {
			logx.Error(err)
		}
		index++
	}

	index = 0
	for index < 10 {
		var v, err = cache.Get([]byte(strconv.Itoa(index)))
		if err != nil {
			logx.Error(err)
			index++
			continue
		}
		logx.Info(string(v))
		index++
	}

	cache.Get([]byte("abccc"))

	logx.Info(cache.HitRate())
}

func TestLocalCache_SetWithExpire(t *testing.T) {
	var cache = NewLocalCache()
	var k = []byte("abc")
	if err := cache.SetWithExpire(k, []byte("123"), 6); err != nil {
		logx.Error(err)
	}

	time.Sleep(2 * time.Second)

	{
		var left, err = cache.TTL(k)
		if err != nil {
			logx.Error(err)
		}
		logx.Info(left)
	}

	time.Sleep(5 * time.Second)

	{
		var v, err = cache.Get(k)
		if err != nil {
			logx.Error(err)
		}


		logx.Info(string(v))
	}

}