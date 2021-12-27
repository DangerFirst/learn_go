package main

import (
	"fmt"
)

func main() {
	//converType()
	var sex string
	fmt.Print("请输入性别：")
	fmt.Scanln(&sex)
	fmt.Println(sexJudge(sex))
	fmt.Println("Finish")
}

//
//func converType() {
//	defer func() {
//		if r:=recover();r!=nil{
//			fmt.Println("panic了：",r)
//			debug.PrintStack()
//		}
//	}()
//	var a interface{}
//	a="string a"
//	b:=a.(int)
//	fmt.Println(b)
//}

func sexJudge(sex string) (sexWight int, err error) {
	if sex != "男" && sex != "女" {
		return 0, fmt.Errorf("输入的性别有误")
	} else if sex == "男" {
		return 1, nil
	}
	return 0, nil
}
