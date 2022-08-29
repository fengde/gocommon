package funcx

import "time"

// 执行函数，带重试，指定重试间的时间间隔
// 如果f返回error != nil，将触发重试
// 指定retry是重试的次数，最终执行的次数是 retry+1
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
