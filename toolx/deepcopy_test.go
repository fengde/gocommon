package toolx

import (
	"testing"
)

func TestDeepCopy(t *testing.T) {
	// slice 非引用检测
	src1 := []string{"a", "b", "c"}
	r1 := DeepCopy(src1).([]string)
	t.Log(r1)
	src1[0] = "aa"
	t.Log(r1)
	// slice 引用检测
	src2 := []map[string]int{
		map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		},
	}
	r2 := DeepCopy(src2).([]map[string]int)
	t.Log(r2)
	src2[0]["a"] = 2
	t.Log(r2)
	src2 = append(src2, map[string]int{
		"e": 4,
		"f": 5,
	})
	t.Log(r2)
	// struct 检测
	type group struct {
		People []string
		name   []string
	}
	src3 := group{
		People: []string{"zs", "ls", "ww"},
		name:   []string{"zs", "ls", "ww"},
	}
	r3 := DeepCopy(src3).(group)
	t.Log(r3)
	src3.People = append(src3.People, "fe")
	t.Log(r3)
}
