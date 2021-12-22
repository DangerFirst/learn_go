package input

import "fmt"

func MaterialFromInput() (name, sex string, tall, weight float64, age, sexWeight int) {
	fmt.Scanln(&name)
	//判断年龄是否满18岁
	for {
		fmt.Print("输入年龄：")
		fmt.Scanln(&age)
		if age >= 18 {
			break
		} else {
			fmt.Print("未满18无法评估，请重新")
		}
	}
	//判断输入性别是否正确
	for {
		fmt.Print("输入性别：")
		fmt.Scanln(&sex)
		//性别转性别比重
		if sex == "男" {
			sexWeight = 1
			break
		} else if sex == "女" {
			sexWeight = 0
			break
		} else {
			fmt.Print("性别错误，请重新")
		}
	}
	fmt.Print("输入身高（厘米）：")
	fmt.Scanln(&tall)
	fmt.Print("输入体重（公斤）：")
	fmt.Scanln(&weight)
	return
}