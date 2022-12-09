package mysqlx

import (
	"fmt"
	"strings"
	"time"

	"github.com/fengde/gocommon/timex"
)

// Datetime 返回正常的时间格式 "2006-01-02 15:04:05"
type Datetime time.Time

func (t Datetime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", t.Format())
	return []byte(stamp), nil
}

func (t *Datetime) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	src := strings.Trim(string(b), `"`)
	_t, err := timex.String2Time(src)
	if err != nil {
		return err
	}

	*t = Datetime(_t)
	return nil
}

func (t Datetime) Time() time.Time {
	return time.Time(t)
}

func (t Datetime) Format() string {
	return t.Time().Format("2006-01-02 15:04:05")
}

// Date 返回正常的时间格式 "2006-01-02"
type Date time.Time

func (t Date) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", t.Format())
	return []byte(stamp), nil
}

func (t *Date) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	src := strings.Trim(string(b), `"`)
	_t, err := timex.DateString2Time(src)
	if err != nil {
		return err
	}
	*t = Date(_t)
	return nil
}

func (t Date) Time() time.Time {
	return time.Time(t)
}

func (t Date) Format() string {
	return t.Time().Format("2006-01-02")
}


type NormalTime Datetime
type NormalDate Date