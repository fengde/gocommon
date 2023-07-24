package namex

import "testing"

func TestGenerateName2Word(t *testing.T) {
	for range make([]struct{}, 10) {
		t.Log(GenerateName2Word())
		t.Log(GenerateName3Word())
		t.Log("冯" + GetLastName() + GetLastName())
	}
}
