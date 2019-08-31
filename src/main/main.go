package main

import (
	"fmt"
	. "github.com/lxn/walk/declarative"
	"os"
)
import "../window"

/*var helpDir = func() {
	fmt.Println("====================================================")
	fmt.Println("mkdir.exe用法：")
	fmt.Println("mkdir.exe path excelPath")
	fmt.Println("如：mkdir.exe F:/project F:/doc/目录结构.xlsx")
	fmt.Println("====================================================")
}*/

func main() {
	sw := &window.SelectWindow{}
	MW := MainWindow{
		AssignTo: &sw.MainWindow,
		//Icon:     "test.ico",
		Title:    "选择excel和目录",
		MinSize:  Size{300, 100},
		Size:     Size{400, 250},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				MaxSize: Size{0, 50},
				Layout:  HBox{},
				Children: []Widget{
					Label{Text: "excel："},
					LineEdit{AssignTo: &sw.ExcelFile},
					PushButton{
						OnClicked: sw.SelectExcel,
						Text:     "浏览...",
					},
				},
			},
			Composite{
				MaxSize: Size{0, 50},
				Layout:  HBox{},
				Children: []Widget{
					Label{Text: "目录："},
					LineEdit{AssignTo: &sw.Dir},
					PushButton{
						OnClicked: sw.SelectDir,
						Text:     "浏览...",
					},
				},
			},
			Composite{
				MaxSize: Size{0, 50},
				Layout:  HBox{},
				Children: []Widget{
					PushButton{
						OnClicked: sw.Confirm,
						Text:     "确定",
					},
					PushButton{
						OnClicked: sw.Cancel,
						Text:     "取消",
					},
				},
			},

		},
	}
	if _, err := MW.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
