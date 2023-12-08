package setx

import (
	"testing"
)

func TestSet_Delete(t *testing.T) {
	var s Set
	v := 1
	s.Store(v)
	if s.Has(v) {
		t.Log("store has true")
	}
	s.Delete(v)
	if s.Has(v) {
		t.Log("delete has true")
	}
	t.Log(s.Items())
}

func TestSet_Has(t *testing.T) {
	var s Set
	v := 1
	if s.Has(v) {
		t.Log("has true")
	}
	s.Store(v)
	if s.Has(v) {
		t.Log("has true 2")
	}
	s.Delete(v)
	if s.Has(v) {
		t.Log("has true 3")
	}
}

func TestSet_Items(t *testing.T) {
	var s Set
	s.Store(1, 2, 3)
	t.Log(s.Items())
}

func TestSet_Store(t *testing.T) {
	var s Set
	s.Store(1)
	s.Store(1)
	s.Store(2)
	t.Log(s.Items())
}

func TestSet_LeftDifference(t *testing.T) {
	set1 := Set{}
	set2 := Set{}
	set1.Store(1, 2, 3, 4)
	set2.Store(1, 3, 5, 8)
	t.Log(set1.LeftDifference(&set2))
	t.Log(set1.RightDifference(&set2))
	t.Log(set1.InnerHave(&set2))
}
