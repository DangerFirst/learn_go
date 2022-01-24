package main

import "fmt"

func GetMaterialFromInput() *Person {
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

	return &Person{
		Name:   name,
		Sex:    sex,
		Age:    age,
		Tall:   tall,
		Weight: weight,
	}
}

func getFakePersonInfo() *Person {
	return &Person{
		Name:    "小黑",
		Sex:     "男",
		Age:     19,
		Tall:    1.8,
		Weight:  65,
		Bmi:     20,
		FatRate: 0.24,
	}
}
