package mysqlx

import (
	"testing"
	"time"

	"github.com/pkg/errors"

	"github.com/fengde/gocommon/logx"
)

func TestNewCluster(t *testing.T) {
	db, err := NewCluster([]string{""}, time.Minute, 10, 10)
	if err != nil {
		logx.Error(err)
		return
	}

	row1, err1 := db.Insert("user", map[string]interface{}{
		"age": 10,
	})
	logx.Info(row1, err1)

	var unit struct {
		Id int64      `xorm:"id"`
		T1 NormalDate `xorm:"t1"`
		T2 NormalTime `xorm:"t2"`
		T3 string     `xorm:"t3"`
	}
	exist4, err4 := db.QueryOne(`select id from user where age=10 limit 1`, nil, &unit)
	logx.Info(unit, exist4, err4)

	var id int64
	exist5, err5 := db.QueryOne(`select id from user where age=10 limit 1`, nil, &id)
	logx.Info(id, exist5, err5)

	effect, err6 := db.Update(`user`, map[string]interface{}{
		"age": 1006,
	}, map[string]interface{}{
		"age": 10,
	})
	logx.Info(effect, err6)

	row2, err2 := db.Exec(`delete from user where age=?`, 18)
	logx.Info(row2, err2)

	row3, err3 := db.Delete("user", map[string]interface{}{
		"age": 10,
	})
	logx.Info(row3, err3)

	db.DoTransaction(func(session *Session) error {
		lastid, err := session.Insert(`user`, map[string]interface{}{
			"age": 1002,
		})
		logx.Info(lastid, err)

		var id int64
		exist, err := session.QueryOne(`select id from user where age=1002`, nil, &id)
		logx.Info(id, exist, err)

		return errors.New("i want rollback")
	})
}
