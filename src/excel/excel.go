package excel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func Read(excelPath string) [][][]string {
	result := [][][]string{}
	xlsx, err := excelize.OpenFile(excelPath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	sheetMap := xlsx.GetSheetMap()
	for _, sheetName := range sheetMap {
		//fmt.Printf("sheetName:%s\n", sheetName)
		rows := xlsx.GetRows(sheetName)
		/*for i, row := range rows {
			fmt.Printf("row index:%d\n", i)
			for j, colCell := range row {
				fmt.Printf("col index:%d,colCell:%s\n", j, colCell)
			}
		}*/
		result = append(result, rows)
	}
	return result
}
