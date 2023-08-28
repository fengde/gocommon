package timex

import (
	"time"
)

const (
	DATE_LAYOUT     = "2006-01-02"
	DATETIME_LAYOUT = "2006-01-02 15:04:05"
)

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
	t, err := time.ParseInLocation(DATETIME_LAYOUT, s, loc)
	if err != nil {
		return t, err
	}
	return t, nil
}

// DateString2Time 字符串日期转Time对象
func DateString2Time(s string) (time.Time, error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, err := time.ParseInLocation(DATE_LAYOUT, s, loc)
	if err != nil {
		return t, err
	}
	return t, nil
}

// Time2String Time对象转字符串时间
func Time2String(t time.Time) string {
	return t.Format(DATETIME_LAYOUT)
}

// Time2Unix Time对象转时间戳
func Time2Unix(t time.Time) int64 {
	return t.Unix()
}

// NowTimeString 返回当前字符串时间格式，如 2022-02-10 00:00:00
func NowTimeString() string {
	return time.Now().Format(DATETIME_LAYOUT)
}

// NowDateString 返回当前字符串日期格式，如 2022-02-10
func NowDateString() string {
	return time.Now().Format(DATE_LAYOUT)
}

// NowUnix 返回当前的unix时间戳
func NowUnix() int64 {
	return time.Now().Unix()
}

// NowUnixMilli 返回当前的unix 毫秒
func NowUnixMilli() int64 {
	return NowUnixNano() / 1000000
}

// NowUnixMicro 返回当前的unix 微秒
func NowUnixMicro() int64 {
	return NowUnixNano() / 1000
}

// NowUnixNano 返回当前的unix 纳秒
func NowUnixNano() int64 {
	return time.Now().UnixNano()
}

// Now 返回当前时间
func Now() time.Time {
	return time.Now()
}

// GetTodayStartTime 获取今天开始时间，如 "xxxx-xx-xx 00:00:00"
func GetTodayStartTime() string {
	return NowDateString() + " 00:00:00"
}

// GetTodayEndTime 获取今天结束时间，如 "xxxx-xx-xx 23:59:59"
func GetTodayEndTime() string {
	return NowDateString() + " 23:59:59"
}

// AddSecond 增加秒
func AddSecond(t time.Time, second int64) time.Time {
	return t.Add(time.Second * time.Duration(second))
}

// AddMinute 增加分钟
func AddMinute(t time.Time, minute int64) time.Time {
	return t.Add(time.Minute * time.Duration(minute))
}

// AddHour 增加小时
func AddHour(t time.Time, hour int64) time.Time {
	return t.Add(time.Hour * time.Duration(hour))
}

// AddDate 新增年月日，year，month, day 均可以指定负数，代表之前的某个时间
func AddDate(t time.Time, year, month, day int) time.Time {
	return t.AddDate(year, month, day)
}
