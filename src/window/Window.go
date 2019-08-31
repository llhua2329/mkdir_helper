package window

import (
	"fmt"
	"github.com/lxn/walk"
	"os"
	"os/exec"
	"strings"
)

import "../excel"
import "../mkdir"

type SelectWindow struct {
	*walk.MainWindow
	ExcelFile *walk.LineEdit
	Dir *walk.LineEdit
	path string
}
func (win *SelectWindow) SelectExcel() {
	dlg := new(walk.FileDialog)
	dlg.FilePath = win.path
	dlg.Title = "选择excel文件"
	dlg.Filter = "excel文件(*.xlsx)|*.xlsx"

	if ok, err := dlg.ShowOpen(win); err != nil {
		//mw.excelFile.AppendText("Error : File Open\r\n")
		win.ExcelFile.SetText("错误 : 文件打开出错\r\n")
		return
	} else if !ok {
		//mw.excelFile.SetText("取消\r\n")
		return
	}
	win.path = dlg.FilePath
	//s := fmt.Sprintf("Select : %s\r\n", mw.path)
	win.ExcelFile.SetText(dlg.FilePath)
}
func (win *SelectWindow) SelectDir(){
	dlg1 := new(walk.FileDialog)
	dlg1.FilePath = win.path
	dlg1.Title = "选择要创建在哪个文件夹下面"
	if ok, err := dlg1.ShowBrowseFolder(win); err != nil {
		//mw.excelFile.AppendText("Error : File Open\r\n")
		win.Dir.SetText("错误 : 文件夹选择出错\r\n")
		return
	} else if !ok {
		//mw.excelFile.SetText("取消\r\n")
		return
	}
	win.path = dlg1.FilePath
	//s := fmt.Sprintf("Select : %s\r\n", mw.path)
	win.Dir.SetText(dlg1.FilePath)
}

func (win *SelectWindow) Confirm(){
	excelPath := win.ExcelFile.Text()
	if excelPath == ""{
		walk.MsgBox(win, "excel文件提示", "excel文件路径不能为空，请重新选择！", walk.MsgBoxIconWarning)
		return
	}
	if _, err := os.Stat(excelPath); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			walk.MsgBox(win, "excel文件提示", "该路径下excel文件不存在，请重新选择！", walk.MsgBoxIconWarning)
			return
		} else {
			// other error
			walk.MsgBox(win, "excel文件提示", "excel文件判断出错，请重新选择！", walk.MsgBoxIconWarning)
			return
		}
	}
	path := win.Dir.Text()
	if path == "" {
		path = getCurrentPath()
	}
	//执行读取excel并创建目录
	createDir(path, excelPath)

	win.Close()
}
func (win *SelectWindow) Cancel(){
	win.Close()
}

func createDir(path string, excelPath string) {
	fmt.Printf("创建路径：%s\n", path)
	fmt.Printf("参数excel：%s\n", excelPath)
	sheets := excel.Read(excelPath)
	if !(strings.HasSuffix(path, "/") || strings.HasSuffix(path, "\\")) {
		path += "/"
	}
	for _, sheet := range sheets {
		mkdir.MkdirHelper(path, sheet)
	}
}

func getCurrentPath() string {
	s, _ := exec.LookPath(os.Args[0])
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}