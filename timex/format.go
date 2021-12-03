package timex

import (
	"time"
)

var layout = "2006-01-02 15:04:05"


// ParseTime 请使用String2Time
func ParseTime(s string) (time.Time, error) {
	return String2Time(s)
}

// Unix2String 时间戳转字符串时间
func Unix2String(unix int64) string {
	return Time2String(time.Unix(unix, 0))
}

// Unix2Time 时间戳转Time对象
func Unix2Time(unix int64) time.Time {
	return time.Unix(unix, 0)
}

// String2Unix 字符串时间转时间戳
func String2Unix(s string) int64 {
	t, err := ParseTime(s)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// String2Time 字符串时间转Time对象
func String2Time(s string) (time.Time, error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, err := time.ParseInLocation(layout, s, loc)
	if err != nil {
		return t, err
	}
	return t, nil
}

// Time2String Time对象转字符串时间
func Time2String(t time.Time) string {
	return t.Format(layout)
}

// Time2Unix Time对象转时间戳
func Time2Unix(t time.Time) int64 {
	return t.Unix()
}
