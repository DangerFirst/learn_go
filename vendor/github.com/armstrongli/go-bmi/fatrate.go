package gobmi

import "fmt"

func Fatrate(age int, sex string, bmi float64) (fatRate float64, err error) {
	if age < 0 {
		err = fmt.Errorf("age cannot be negative")
		return
	}
	if age == 0 {
		err = fmt.Errorf("age cannot be 0")
		return
	}
	if age >= 150 {
		err = fmt.Errorf("age cannot exceed 150")
		return
	}
	if bmi < 0 {
		err = fmt.Errorf("bmi cannot be negative")
		return
	}
	if bmi == 0 {
		err = fmt.Errorf("bmi cannot be 0")
		return
	}
	if sex != "男" && sex != "女" {
		err = fmt.Errorf("sex cannot be different from male and female")
		return
	}
	sexWeight := 0
	if sex == "男" {
		sexWeight = 1
	}
	//计算体脂率
	fatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	return
}
