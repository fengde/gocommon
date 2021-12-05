package mysqlx

import (
	"fmt"
	"time"
)

// NormalTime 返回正常的时间格式 "2006-01-02 15:04:05"
type NormalTime time.Time

func (t NormalTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func (t NormalTime) Time() time.Time {
	return time.Time(t)
}

func (t NormalTime) Format() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}

// NormalDate 返回正常的时间格式 "2006-01-02"
type NormalDate time.Time

func (t NormalDate) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02"))
	return []byte(stamp), nil
}

func (t NormalDate) Time() time.Time {
	return time.Time(t)
}

func (t NormalDate) Format() string {
	return time.Time(t).Format("2006-01-02")
}