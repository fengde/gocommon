package timex

import (
	"testing"
	"time"
)

func TestAll(t *testing.T) {
	t.Log(Time2Unix(Unix2Time(1)))
	t.Log(Time2String(time.Now()))
	t1, _ := String2Time("2020-01-01 00:00:00")
	t.Log(Time2String(t1))
	t.Log(Time2Unix(t1))
}
