package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filePath := "E:/go语言/go/go_file/荣霸天.self.information"
	writeFile(filePath)
	readFile(filePath)
}

func readFile(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	infos := strings.Split(string(data), ",")
	fmt.Println("开始计算体脂信息：", infos)
}

func writeFile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息为：", err)
		os.Exit(1)
	}
	defer file.Close()
	_, err = file.Write([]byte("小强,男,1.70,...."))
	fmt.Println(err)
	//handle error
}
