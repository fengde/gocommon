package slicex

// Contains 判断是否包含
func Contains[T comparable](arr []T, item T) bool {
	for _, t := range arr {
		if t == item {
			return true
		}
	}
	return false
}

// RemoveRepeat 删除重复元素
func RemoveRepeat[T comparable](arr []T) []T {
	var newArr []T
	var m = map[T]struct{}{}
	for _, item := range arr {
		if _, ok := m[item]; !ok {
			newArr = append(newArr, item)
			m[item] = struct{}{}
		}
	}

	return newArr
}

// ForEach 遍历数组中的元素
func ForEach[T comparable](arr []T, fn func(t T)) {
	for _, t := range arr {
		fn(t)
	}
}
