package mysqlx

import (
	"fmt"
	"strings"
	"time"

	"github.com/fengde/gocommon/timex"
)

// NormalTime 返回正常的时间格式 "2006-01-02 15:04:05"
type NormalTime time.Time

func (t NormalTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", t.Format())
	return []byte(stamp), nil
}

func (t *NormalTime) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	src := strings.Trim(string(b), `"`)
	_t, err := timex.String2Time(src)
	if err != nil {
		return err
	}

	*t = NormalTime(_t)
	return nil
}

func (t NormalTime) Time() time.Time {
	return time.Time(t)
}

func (t NormalTime) Format() string {
	return t.Time().Format("2006-01-02 15:04:05")
}

// NormalDate 返回正常的时间格式 "2006-01-02"
type NormalDate time.Time

func (t NormalDate) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", t.Format())
	return []byte(stamp), nil
}

func (t *NormalDate) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	src := strings.Trim(string(b), `"`)
	_t, err := timex.DateString2Time(src)
	if err != nil {
		return err
	}
	*t = NormalDate(_t)
	return nil
}

func (t NormalDate) Time() time.Time {
	return time.Time(t)
}

func (t NormalDate) Format() string {
	return t.Time().Format("2006-01-02")
}
