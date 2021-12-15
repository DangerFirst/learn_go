package main

import (
	"fmt"
	"strings"
)

func main() {
	//定义要转换的数组
	slice := []string{"Hello小明", "小case"}
	//调用转换函数
	transformation(slice)
	//输出转换后的数组
	fmt.Println(slice)
}

func transformation(beforTrans []string) (afferTrans string) {
	for _, sliceContent := range beforTrans {
		for _, cutString := range strings.Split(sliceContent, "") {
			fmt.Println(cutString)
		}
	}
	return
}
