package mysqlx

import (
	"testing"
	"time"

	"github.com/fengde/gocommon/jsonx"
	"github.com/fengde/gocommon/logx"
)

func TestNormalTime_MarshalJSON(t *testing.T) {
	// var t1 struct {
	// 	Time NormalDate
	// }
	// t1.Time = NormalDate(time.Now())
	// s := jsonx.MarshalToStringNoErr(t1)
	// t.Log(s)
	// if err := jsonx.UnmarshalString(s, &t1); err != nil {
	// 	t.Log(err)
	// }
	// t.Log(t1.Time.Format())

	db, err := NewCluster([]string{""}, time.Minute, 10, 10)
	if err != nil {
		logx.Error(err)
		return
	}

	var unit struct {
		Id int64      `xorm:"id"`
		T1 NormalDate `xorm:"t1"`
		T2 NormalTime `xorm:"t2"`
		T3 string     `xorm:"t3"`
	}
	exist4, err4 := db.QueryOne(`select id, t1, t2, t3 from user where id=11`, nil, &unit)
	logx.Info(jsonx.MarshalToStringNoErr(unit), exist4, err4)

	tmp := jsonx.MarshalToStringNoErr(unit)

	if err := jsonx.UnmarshalString(tmp, &unit); err != nil {
		logx.Error(err)
	}

	logx.Info(jsonx.MarshalToStringNoErr(unit))
}
