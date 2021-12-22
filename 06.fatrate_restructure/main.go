package main

import (
	"fmt"
	"learn.go/06.fatrate_restructure/calc"
	"learn.go/06.fatrate_restructure/input"
	"learn.go/06.fatrate_restructure/suggest"
	"learn.go/06.fatrate_restructure/weithercontinul"
)

func main() {
	for {
		//定义变量
		var names, sexs [999]string
		var talls, weights, bmis, fatRates [999]float64
		var ages, sexWeights [999]int
		var num int
		var fatRateSum float64
		//输入人数
		fmt.Print("输入人数：")
		fmt.Scanln(&num)
		//获取所有人的参数
		for i := 0; i < num; i++ {
			fmt.Printf("输入第%v个人的姓名：", i+1)
			names[i], sexs[i], talls[i], weights[i], ages[i], sexWeights[i] = input.MaterialFromInput()
			bmis[i]=calc.CalcuBmi(weights[i], talls[i])
			fatRates[i] = calc.CalcuFatrate(ages[i], sexWeights[i],bmis[i])
			fatRateSum += fatRates[i]
		}
		//判断体脂情况
		for j := 0; j < num; j++ {
			suggest.Suggest(names[j], sexs[j], ages[j], bmis[j], fatRates[j])
		}
		//计算平均体脂率
		calc.AvgFatRate(num, fatRateSum)
		//是否继续
		if weithercontinul.WeitherContinul() {
			return
		}
	}
}






