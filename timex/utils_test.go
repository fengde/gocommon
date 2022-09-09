package timex

import (
	"testing"
	"time"
)

func TestRecentDateByWeekday(t *testing.T) {
	before, after := RecentDateByWeekday(time.Now().Weekday().String())
	t.Log(before, after)

	before, after = RecentDateByWeekday("Sunday")
	t.Log(before, after)
}

func TestRecentDatetimeByWeektime(t *testing.T) {
	before, after := RecentDatetimeByWeektime("Sunday 20:00:01")
	t.Log(before, after)
	before, after = RecentDatetimeByWeektime("Friday 20:00:01")
	t.Log(before, after)

	before, after = RecentDatetimeByWeektime("Friday 10:00:01")
	t.Log(before, after)
}
