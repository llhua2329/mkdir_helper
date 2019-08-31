# mkdir_helper（创建目录助手类程序）

### 使用
* 目录结构在excel文件里定义好，运行后。
  * 1.“excel”输入框里，浏览选择定义目录结构的excel文件。
  * 2.“目录”输入框里，浏览要创建目录的父目录。
  * 3.点“确定”运行。

### 编译
* 命令
`
go build -ldflags="-H windowsgui" main.go
`
