package excelx

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/fengde/gocommon/logx"
	"strconv"
)

// CreateExcelWithMergeCell 合并单元格
func CreateExcelWithMergeCell(filepath string, titles []string, rows [][]interface{}, needMergeTitles []string) error {
	f := excelize.NewFile()
	// 创建一个工作表
	const Sheet1 = "Sheet1"
	index := f.NewSheet(Sheet1)

	var titleIndex2Axis = map[int]string{}
	var titleName2Axies = map[string]string{}
	var titleName2Index = map[string]int{}

	for i, title := range titles {
		axisPrefix := string(rune(65+i))
		f.SetCellValue(Sheet1, axisPrefix+ "1", title)
		titleIndex2Axis[i] = axisPrefix
		titleName2Axies[title] = axisPrefix
		titleName2Index[title] = i
	}

	for i, row := range rows {
		for j, col := range row {
			f.SetCellValue(Sheet1, titleIndex2Axis[j] + strconv.Itoa(i+2), col)
		}
	}

	var lastIndexs []int

	for _, title := range needMergeTitles {
		axisPrefix := titleName2Axies[title]

		lastIndexs = append(lastIndexs, titleName2Index[title])

		var hcellRowID, vcellRowID = 2, 2

		for i := range rows {
			var tryMerge bool
			if i < len(rows) - 1 {
				var allHit = true
				for _, index := range lastIndexs {
					if rows[i][index] == rows[i+1][index] {
						continue
					} else {
						allHit = false
						break
					}
				}
				if allHit {
					vcellRowID++
				} else {
					tryMerge = true
				}
			} else {
				tryMerge = true
			}


			if tryMerge && vcellRowID > hcellRowID {
				logx.Info(fmt.Sprintf("%v%v", axisPrefix, hcellRowID), " ", fmt.Sprintf("%v%v", axisPrefix, vcellRowID))
				f.MergeCell(Sheet1, fmt.Sprintf("%v%v", axisPrefix, hcellRowID), fmt.Sprintf("%v%v", axisPrefix, vcellRowID))
				hcellRowID = hcellRowID
			}

			if  !tryMerge && vcellRowID == hcellRowID {
				hcellRowID = vcellRowID + i + 1
				vcellRowID = hcellRowID
			}
		}
	}

	f.SetActiveSheet(index)

	f.SaveAs(filepath)


	return nil
}