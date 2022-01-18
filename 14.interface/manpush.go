package main

import "fmt"

type manLegend struct {
}

func (*manLegend) OpenTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("人工打开冰箱")
	return nil
}

func (*manLegend) PutElephantIntoRefrigerator(Elephant, Refrigerator) error {
	fmt.Println("人工装进冰箱")
	return nil
}

func (*manLegend) CloseTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("人工关闭冰箱")
	return nil
}
