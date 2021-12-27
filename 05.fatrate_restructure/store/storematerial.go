package store

import (
	"fmt"
	"learn.go/05.fatrate_restructure/calc"
	"learn.go/05.fatrate_restructure/input"
)

func StoreMaterial(num int) (names, sexs [999]string, bmis, fatRates [999]float64, ages [999]int, fatRateSum float64) {
	var talls, weights [999]float64
	var sexWeights [999]int
	//获取所有人的参数
	for i := 0; i < num; i++ {
		fmt.Printf("输入第%v个人的姓名：", i+1)
		names[i], sexs[i], talls[i], weights[i], ages[i], sexWeights[i] = input.MaterialFromInput()
		bmis[i] = calc.CalcuBmi(weights[i], talls[i])
		fatRates[i] = calc.CalcuFatrate(ages[i], sexWeights[i], bmis[i])
		fatRateSum += fatRates[i]
	}
	return
}
