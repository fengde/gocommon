package toolx

// Repeat 重复执行函数
func Repeat(count int, f func()) {
	for i := 0; i < count; i++ {
		f()
	}
}
