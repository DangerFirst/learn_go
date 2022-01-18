package main

func getFakePersonInfo() *Person {
	return &Person{
		Name:    "小黑",
		Sex:     "男",
		Age:     19,
		Tall:    1.8,
		Weight:  65,
		Bmi:     20,
		FatRate: 0.24,
	}
}
