package timex

import "strings"

const (
	Monday    = "Monday"    // 星期一
	Tuesday   = "Tuesday"   // 星期二
	Wednesday = "Wednesday" // 星期三
	Thursday  = "Thursday"  // 星期四
	Friday    = "Friday"    // 星期五
	Saturday  = "Saturday"  // 星期六
	Sunday    = "Sunday"    // 星期日
)

// 例子：weekday=Monday，返回当前时间，前一个Mondy的日期 和 后一个Monday的日期
func RecentDateByWeekday(weekday string) (before string, after string) {
	var m = map[string]int{
		Monday:    1,
		Tuesday:   2,
		Wednesday: 3,
		Thursday:  4,
		Friday:    5,
		Saturday:  6,
		Sunday:    7,
	}

	now := Now()
	num1 := m[now.Weekday().String()]
	num2 := m[weekday]

	switch {
	case num1 > num2:
		before = now.AddDate(0, 0, num2-num1).Format(DATE_LAYOUT)
		after = now.AddDate(0, 0, num2+7-num1).Format(DATE_LAYOUT)
	case num1 == num2:
		before = now.Format(DATE_LAYOUT)
		after = now.AddDate(0, 0, 7).Format((DATE_LAYOUT))
	case num1 < num2:
		before = now.AddDate(0, 0, num2-num1-7).Format(DATE_LAYOUT)
		after = now.AddDate(0, 0, num2-num1).Format(DATE_LAYOUT)
	}
	return
}

// 例子：weektime="Monday 20:00:00"，返回当前时间，前一个Mondy的时间 和 后一个Monday时间
func RecentDatetimeByWeektime(weektime string) (before string, after string) {
	var m = map[string]int{
		Monday:    1,
		Tuesday:   2,
		Wednesday: 3,
		Thursday:  4,
		Friday:    5,
		Saturday:  6,
		Sunday:    7,
	}

	items := strings.Split(weektime, " ")
	if len(items) == 2 {
		weekday := items[0]
		dayTime := items[1]

		now := Now()
		num1 := m[now.Weekday().String()]
		num2 := m[weekday]

		switch {
		case num1 > num2:
			before = now.AddDate(0, 0, num2-num1).Format(DATE_LAYOUT) + " " + dayTime
			after = now.AddDate(0, 0, num2+7-num1).Format(DATE_LAYOUT) + " " + dayTime
		case num1 == num2:
			t := now.Format(DATE_LAYOUT) + " " + dayTime
			if String2Unix(t) < now.Unix() {
				before = t
				t2, _ := String2Time(t)
				after = Time2String(t2.AddDate(0, 0, 7))
			} else {
				after = t
				t2, _ := String2Time(t)
				before = Time2String(t2.AddDate(0, 0, -7))
			}
		case num1 < num2:
			before = now.AddDate(0, 0, num2-num1-7).Format(DATE_LAYOUT) + " " + dayTime
			after = now.AddDate(0, 0, num2-num1).Format(DATE_LAYOUT) + " " + dayTime
		}
	}

	return
}

// StringTimeCompareIsT1BeforeT2 比较两个时间(2006-01-02 15:04:05)，如果t1 在 t2 之前返回true，否则返回false
func StringTimeCompareIsT1BeforeT2(t1 string, t2 string) bool {
	return String2Unix(t1) < String2Unix(t2)
}

// StringTimeCompareIsT1AfterT2 比较两个时间(2006-01-02 15:04:05)，如果t1 在 t2 之后返回true，否则返回false
func StringTimeCompareIsT1AfterT2(t1 string, t2 string) bool {
	return String2Unix(t1) > String2Unix(t2)
}
