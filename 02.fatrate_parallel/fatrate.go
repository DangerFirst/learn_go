package main

import (
	"fmt"
)

func main() {
	for {
		//定义变量
		var names, sexs [999]string
		var talls, weights, bmis, fatRates [999]float64
		var ages [999]int
		var reEnter string
		var sexWeight, num int
		var fatRateSum float64
		//输入参数
		fmt.Print("输入人数：")
		fmt.Scanln(&num)
		for i := 0; i < num; i++ {
			fmt.Printf("输入第%v个人的姓名：", i+1)
			fmt.Scanln(&names[i])
			//判断年龄是否满18岁
			for {
				fmt.Print("输入年龄：")
				fmt.Scanln(&ages[i])
				if ages[i] >= 18 {
					break
				} else {
					fmt.Print("未满18无法评估，请重新")
				}
			}
			//判断输入性别是否正确
			for {
				fmt.Print("输入性别：")
				fmt.Scanln(&sexs[i])
				//性别转性别比重
				if sexs[i] == "男" {
					sexWeight = 1
					break
				} else if sexs[i] == "女" {
					sexWeight = 0
					break
				} else {
					fmt.Print("性别错误，请重新")
				}
			}
			fmt.Print("输入身高（厘米）：")
			fmt.Scanln(&talls[i])
			fmt.Print("输入体重（公斤）：")
			fmt.Scanln(&weights[i])
			//计算体脂率
			bmis[i] = weights[i] / ((talls[i] * talls[i]) / 10000)
			fatRates[i] = (1.2*bmis[i] + 0.23*float64(ages[i]) - 5.4 - 10.8*float64(sexWeight)) / 100
			fatRateSum += fatRates[i]
		}
		//判断体脂情况
		for j := 0; j < num; j++ {
			fmt.Println("姓名：", names[j])
			fmt.Printf("BMI：%.2f\n", bmis[j])
			fmt.Printf("体脂率：%.2f\n", fatRates[j])
			fmt.Print("评估及建议：")
			//男性体脂判断
			age := ages[j]
			fatRate := fatRates[j]
			if sexs[j] == "男" {
				if age >= 18 && age < 40 {
					if fatRate < 0 {
						fmt.Println("异常，请检查输入内容重新输入")
					} else if fatRate >= 0 && fatRate < 0.1 {
						fmt.Println("偏瘦，需要多吃点")
					} else if fatRate < 0.16 {
						fmt.Println("标准，不错，保持下去")
					} else if fatRate < 0.21 {
						fmt.Println("偏重，注意饮食，加强锻炼")
					} else if fatRate < 0.26 {
						fmt.Println("肥胖，控制饮食，多多锻炼")
					} else {
						fmt.Println("严重肥胖，少吃多动，坚持下去")
					}
				} else if age >= 40 && age < 60 {
					if fatRate < 0 {
						fmt.Println("异常，请检查输入内容重新输入")
					} else if fatRate >= 0 && fatRate < 0.11 {
						fmt.Println("偏瘦，需要多吃点")
					} else if fatRate < 0.17 {
						fmt.Println("标准，不错，保持下去")
					} else if fatRate < 0.22 {
						fmt.Println("偏重，注意饮食，加强锻炼")
					} else if fatRate < 0.27 {
						fmt.Println("肥胖，控制饮食，多多锻炼")
					} else {
						fmt.Println("严重肥胖，少吃多动，坚持下去")
					}
				} else {
					if fatRate < 0 {
						fmt.Println("异常，请检查输入内容重新输入")
					} else if fatRate >= 0 && fatRate < 0.13 {
						fmt.Println("偏瘦，需要多吃点")
					} else if fatRate < 0.19 {
						fmt.Println("标准，不错，保持下去")
					} else if fatRate < 0.24 {
						fmt.Println("偏重，注意饮食，加强锻炼")
					} else if fatRate < 0.29 {
						fmt.Println("肥胖，控制饮食，多多锻炼")
					} else {
						fmt.Println("严重肥胖，少吃多动，坚持下去")
					}
				}
				//女性体脂判断
			} else {
				if age >= 18 && age < 40 {
					if fatRate < 0 {
						fmt.Println("异常，请检查输入内容重新输入")
					} else if fatRate >= 0 && fatRate < 0.2 {
						fmt.Println("偏瘦，需要多吃点")
					} else if fatRate < 0.27 {
						fmt.Println("标准，不错，保持下去")
					} else if fatRate < 0.34 {
						fmt.Println("偏重，注意饮食，加强锻炼")
					} else if fatRate < 0.39 {
						fmt.Println("肥胖，控制饮食，多多锻炼")
					} else {
						fmt.Println("严重肥胖，少吃多动，坚持下去")
					}
				} else if age >= 40 && age < 60 {
					if fatRate < 0 {
						fmt.Println("异常，请检查输入内容重新输入")
					} else if fatRate >= 0 && fatRate < 0.21 {
						fmt.Println("偏瘦，需要多吃点")
					} else if fatRate < 0.28 {
						fmt.Println("标准，不错，保持下去")
					} else if fatRate < 0.35 {
						fmt.Println("偏重，注意饮食，加强锻炼")
					} else if fatRate < 0.4 {
						fmt.Println("肥胖，控制饮食，多多锻炼")
					} else {
						fmt.Println("严重肥胖，少吃多动，坚持下去")
					}
				} else {
					if fatRate < 0 {
						fmt.Println("异常，请检查输入内容重新输入")
					} else if fatRate >= 0 && fatRate < 0.22 {
						fmt.Println("偏瘦，需要多吃点")
					} else if fatRate < 0.29 {
						fmt.Println("标准，不错，保持下去")
					} else if fatRate < 0.36 {
						fmt.Println("偏重，注意饮食，加强锻炼")
					} else if fatRate < 0.41 {
						fmt.Println("肥胖，控制饮食，多多锻炼")
					} else {
						fmt.Println("严重肥胖，少吃多动，坚持下去")
					}
				}
			}
		}
		fmt.Printf("总人数：%d，平均体脂率 %.2f\n", num, fatRateSum/float64(num))
		fmt.Print("是否重新输入（y/n）：")
		fmt.Scanln(&reEnter)
		if reEnter != "y" {
			fmt.Println("程序结束")
			return
		}
	}
}
