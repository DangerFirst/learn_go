package main

import (
	"fmt"
)

func main() {
	for {
		//定义变量
		//var names,sexs ...[]string
		//var talls,weights ...[]float64
		//var ages ...[]int
		var weight, tall, bmi, weightrate float64
		var sex, name, reEnter string
		var age, sexWeight int
		//输入参数
		fmt.Print("输入姓名：")
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
		fmt.Print("输入身高（米）：")
		fmt.Scanln(&tall)
		fmt.Print("输入体重（公斤）：")
		fmt.Scanln(&weight)

		//计算体脂率
		bmi = weight / (tall * tall)
		weightrate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
		fmt.Println("姓名：", name)
		fmt.Println("BMI：", bmi)
		fmt.Println("您的体脂率为：", weightrate)
		//判断体脂情况
		fmt.Print("您的体脂程度为：")
		//男性体脂判断
		if sex == "男" {
			if age >= 18 && age < 40 {
				if weightrate < 0 {
					fmt.Println("异常，请检查输入内容重新输入")
				} else if weightrate >= 0 && weightrate < 0.1 {
					fmt.Println("偏瘦，需要多吃点")
				} else if weightrate < 0.16 {
					fmt.Println("标准，不错，保持下去")
				} else if weightrate < 0.21 {
					fmt.Println("偏重，注意饮食，加强锻炼")
				} else if weightrate < 0.26 {
					fmt.Println("肥胖，控制饮食，多多锻炼")
				} else {
					fmt.Println("严重肥胖，少吃多动，坚持下去")
				}
			} else if age >= 40 && age < 60 {
				if weightrate < 0 {
					fmt.Println("异常，请检查输入内容重新输入")
				} else if weightrate >= 0 && weightrate < 0.11 {
					fmt.Println("偏瘦，需要多吃点")
				} else if weightrate < 0.17 {
					fmt.Println("标准，不错，保持下去")
				} else if weightrate < 0.22 {
					fmt.Println("偏重，注意饮食，加强锻炼")
				} else if weightrate < 0.27 {
					fmt.Println("肥胖，控制饮食，多多锻炼")
				} else {
					fmt.Println("严重肥胖，少吃多动，坚持下去")
				}
			} else {
				if weightrate < 0 {
					fmt.Println("异常，请检查输入内容重新输入")
				} else if weightrate >= 0 && weightrate < 0.13 {
					fmt.Println("偏瘦，需要多吃点")
				} else if weightrate < 0.19 {
					fmt.Println("标准，不错，保持下去")
				} else if weightrate < 0.24 {
					fmt.Println("偏重，注意饮食，加强锻炼")
				} else if weightrate < 0.29 {
					fmt.Println("肥胖，控制饮食，多多锻炼")
				} else {
					fmt.Println("严重肥胖，少吃多动，坚持下去")
				}
			}
			//女性体脂判断
		} else {
			if age >= 18 && age < 40 {
				if weightrate < 0 {
					fmt.Println("异常，请检查输入内容重新输入")
				} else if weightrate >= 0 && weightrate < 0.2 {
					fmt.Println("偏瘦，需要多吃点")
				} else if weightrate < 0.27 {
					fmt.Println("标准，不错，保持下去")
				} else if weightrate < 0.34 {
					fmt.Println("偏重，注意饮食，加强锻炼")
				} else if weightrate < 0.39 {
					fmt.Println("肥胖，控制饮食，多多锻炼")
				} else {
					fmt.Println("严重肥胖，少吃多动，坚持下去")
				}
			} else if age >= 40 && age < 60 {
				if weightrate < 0 {
					fmt.Println("异常，请检查输入内容重新输入")
				} else if weightrate >= 0 && weightrate < 0.21 {
					fmt.Println("偏瘦，需要多吃点")
				} else if weightrate < 0.28 {
					fmt.Println("标准，不错，保持下去")
				} else if weightrate < 0.35 {
					fmt.Println("偏重，注意饮食，加强锻炼")
				} else if weightrate < 0.4 {
					fmt.Println("肥胖，控制饮食，多多锻炼")
				} else {
					fmt.Println("严重肥胖，少吃多动，坚持下去")
				}
			} else {
				if weightrate < 0 {
					fmt.Println("异常，请检查输入内容重新输入")
				} else if weightrate >= 0 && weightrate < 0.22 {
					fmt.Println("偏瘦，需要多吃点")
				} else if weightrate < 0.29 {
					fmt.Println("标准，不错，保持下去")
				} else if weightrate < 0.36 {
					fmt.Println("偏重，注意饮食，加强锻炼")
				} else if weightrate < 0.41 {
					fmt.Println("肥胖，控制饮食，多多锻炼")
				} else {
					fmt.Println("严重肥胖，少吃多动，坚持下去")
				}
			}
		}
		fmt.Print("是否重新输入（y/n）：")
		fmt.Scanln(&reEnter)
		if reEnter != "y" {
			fmt.Println("程序结束")
			return
		}
	}
}
