package weithercontinul

import "fmt"

func WeitherContinul() bool {
	var reEnter string
	fmt.Print("是否重新输入（y/n）：")
	fmt.Scanln(&reEnter)
	if reEnter != "y" {
		fmt.Println("程序结束")
		return true
	}
	return false
}
