package slicex

// Deprecated: Use Contains instead.
func StrContains(arr []string, item string) bool {
	for _, t := range arr {
		if t == item {
			return true
		}
	}
	return false
}

// Deprecated: Use Contains instead.
func IntContains(arr []int, item int) bool {
	for _, t := range arr {
		if t == item {
			return true
		}
	}
	return false
}

// Deprecated: Use Contains instead.
func Int64Contains(arr []int64, item int64) bool {
	for _, t := range arr {
		if t == item {
			return true
		}
	}
	return false
}

// Deprecated: Use RemoveRepeat instead.
func StrRemoveRepeat(arr []string) []string {
	var newArr []string
	var m = map[string]int{}
	for _, item := range arr {
		if _, ok := m[item]; !ok {
			newArr = append(newArr, item)
			m[item] = 1
		}
	}

	return newArr
}

// Deprecated: Use RemoveRepeat instead.
func IntRemoveRepeat(arr []int) []int {
	var newArr []int
	var m = map[int]int{}
	for _, item := range arr {
		if _, ok := m[item]; !ok {
			newArr = append(newArr, item)
			m[item] = 1
		}
	}

	return newArr
}

// Deprecated: Use RemoveRepeat instead.
func Int64RemoveRepeat(arr []int64) []int64 {
	var newArr []int64
	var m = map[int64]int{}
	for _, item := range arr {
		if _, ok := m[item]; !ok {
			newArr = append(newArr, item)
			m[item] = 1
		}
	}

	return newArr
}
