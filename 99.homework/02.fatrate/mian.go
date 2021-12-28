package main

import (
	"learn.go/99.homework/02.fatrate/calc"
	"learn.go/99.homework/02.fatrate/continul"
	"learn.go/99.homework/02.fatrate/input"
	"learn.go/99.homework/02.fatrate/suggest"
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
