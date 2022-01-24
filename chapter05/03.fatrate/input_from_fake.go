package main

type inputFromFake struct{}

func (*inputFromFake) GetInput() Person {
	return Person{
		Name:    "小黑",
		Sex:     "男",
		Age:     20,
		Tall:    1.8,
		Weight:  65,
		Bmi:     20,
		FatRate: 0.24,
	}
}
