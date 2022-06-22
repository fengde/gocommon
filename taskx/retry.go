package taskx

import "time"

func Retry(retry int, interval time.Duration, f func(loop int) error) error {
	for i := 0; ; i++ {
		if err := f(i); err != nil {
			if i < retry {
				if interval > 0 {
					time.Sleep(interval)
				}
			} else {
				return err
			}
		} else {
			return nil
		}
	}
}
