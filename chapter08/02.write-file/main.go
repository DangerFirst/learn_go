package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "E:/go语言/go/go_file/荣霸天.new"
	//writeFile(filePath)
	writeFileWithAppend(filePath)
}

func writeFile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息为：", err)
		os.Exit(1)
	}
	defer file.Close()
	_, err = file.Write([]byte("This is first line--"))
	_, err = file.Write([]byte("This is first line\n"))
	fmt.Println(err)
	//handle error
	_, err = file.WriteString("第二行内容\n")
	fmt.Println(err)
	//handle error
	_, err = file.WriteAt([]byte("CHANGED"), 0)
	//file.Sync()//告诉操作系统直接把内容写到磁盘上
	fmt.Println(err)
	//handle error
}

func writeFileWithAppend(filePath string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777) //linux file permission settings
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息为：", err)
		os.Exit(1)
	}
	defer file.Close()
	_, err = file.Write([]byte("This is first line--"))
	_, err = file.Write([]byte("This is first line\n"))
	fmt.Println(err)
	//handle error
	_, err = file.WriteString("第二行内容\n")
	fmt.Println(err)
	//handle error
	_, err = file.WriteAt([]byte("CHANGED"), 0)
	//file.Sync()//告诉操作系统直接把内容写到磁盘上
	fmt.Println(err)
	//handle error
}
