package main

import (
	"fatrate/calc"
	"fatrate/continul"
	"fatrate/input"
	"fatrate/suggest"
)

func main() {
	for { //获取参数
		person := input.GetMaterialFromInput()
		//计算BMI
		calc.Calc{}.BMI(person)
		//计算体脂率
		calc.Calc{}.FatRate(person)
		//获取体脂建议
		suggest.Suggest(person)
		//是否继续
		if continul.WeitherContinul() {
			return
		}
	}
}
