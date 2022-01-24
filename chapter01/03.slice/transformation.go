package main

import (
	"fmt"
	"strings"
)

func main() {
	//定义要转换的数组
	slice := []string{"Hello小明", "小case"}
	fmt.Println("转换前数组为：", slice)
	//调用转换函数
	transformation(slice)
	//输出转换后的数组
	fmt.Println("转换后数组为：", slice)
}

//转换大小写函数
func transformation(beforTrans []string) (afferTrans string) {
	for contentSit, sliceContent := range beforTrans {
		beforTrans[contentSit] = strings.ToUpper(sliceContent)
	}
	return
}
