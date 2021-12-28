package gobmi

func Fatrate(age int, sex string, bmi float64) (fatRate float64) {
	sexWeight:=0
	if sex == "男" {
		sexWeight = 1
	}
	//计算体脂率
	fatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	return
}
