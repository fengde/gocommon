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

func TestIntervalDays(t *testing.T) {
	type args struct {
		t1 time.Time
		t2 time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01",
			args: args{
				t2: time.Now(),
				t1: time.Date(2023, 7, 20, 23, 20, 9, 0, time.Local),
			},
			want: 2,
		},
		{
			name: "02",
			args: args{
				t2: time.Now(),
				t1: time.Now(),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntervalDays(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("IntervalDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
