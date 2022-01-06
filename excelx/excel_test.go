package excelx

import "testing"

func TestCreateExcelWithMergeCell(t *testing.T) {
	CreateExcelWithMergeCell("./test.xlsx", []string{"标题1", "标题2", "标题3"}, [][]interface{} {
		[]interface{}{"a2", "b2", "c3"},
		[]interface{}{"a2", "b2", "c3"},
		[]interface{}{"a2", "b3", "c3"},
	}, []string{"标题1", "标题2", "标题3"})
}
