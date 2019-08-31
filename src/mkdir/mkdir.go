package mkdir

import (
	"os"
	"strings"
)

func MkdirHelper(path string, dirInfo [][]string) {
	baseDir := []string{}
	startIndex := 0
	for i, row := range dirInfo {
		if i == 0 {
			baseDir = row
		} else {
			for index, v := range row {
				if v != "" {
					startIndex = index
					break
				}
			}
			baseDir = append(baseDir[:startIndex], row[startIndex:]...)
		}
		//fmt.Println(baseDir)
		dir := strings.Join(baseDir, "\\") //strings.Replace(strings.Trim(fmt.Sprint(baseDir), "[]"), " ", "/", -1)
		dir = strings.Replace(dir, "/", " ", -1)
		//fmt.Println(dir)

		os.MkdirAll(path + dir, os.ModePerm)
	}
}
