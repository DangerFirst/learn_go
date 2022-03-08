package main

import (
	"fmt"
	"learn.go/pkg/apis"
)

type inputFromStd struct{}

func (*inputFromStd) GetInput() *apis.PersonalInformation {
	var name string
	fmt.Print("输入姓名：")
	fmt.Scanln(&name)

	var sex string
	fmt.Print("输入性别：")
	fmt.Scanln(&sex)

	var age int
	fmt.Print("输入年龄：")
	fmt.Scanln(&age)

	var tall float64
	fmt.Print("输入身高（米）：")
	fmt.Scanln(&tall)

	var weight float64
	fmt.Print("输入体重（千克）：")
	fmt.Scanln(&weight)

	return &apis.PersonalInformation{
		Name:   name,
		Sex:    sex,
		Age:    age,
		Tall:   tall,
		Weight: weight,
	}
}
