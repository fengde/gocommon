package slicex

func StrContains(arr []string, item string) bool {
	for _, t := range arr {
		if t == item {
			return true
		}
	}
	return false
}

func IntContains(arr []int, item int) bool {
	for _, t := range arr {
		if t == item {
			return true
		}
	}
	return false
}

func Int64Contains(arr []int64, item int64) bool {
	for _, t := range arr {
		if t == item {
			return true
		}
	}
	return false
}
