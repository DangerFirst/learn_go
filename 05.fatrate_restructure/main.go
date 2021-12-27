package main

import (
	"fmt"
	"learn.go/05.fatrate_restructure/calc"
	"learn.go/05.fatrate_restructure/store"
	"learn.go/05.fatrate_restructure/suggest"
	"learn.go/05.fatrate_restructure/weithercontinul"
)

func main() {
	for {
		var num int
		//输入人数
		fmt.Print("输入人数：")
		fmt.Scanln(&num)
		//写入各项参数
		names, sexs, bmis, fatRates, ages, fatRateSum := store.StoreMaterial(num)
		//获取体脂建议
		suggest.Suggest(names, sexs, ages, num, bmis, fatRates)
		//获得平均体脂率
		calc.AvgFatRate(num, fatRateSum)
		//是否继续
		if weithercontinul.WeitherContinul() {
			return
		}
	}
}
