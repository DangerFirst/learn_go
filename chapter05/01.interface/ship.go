package main

import "fmt"

type shipLegend struct {
}

func (*shipLegend) OpenTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("打开船上冰箱")
	return nil
}

func (*shipLegend) PutElephantIntoRefrigerator(Elephant, Refrigerator) error {
	fmt.Println("装进船上冰箱")
	return nil
}

func (*shipLegend) CloseTheDoorOfRefrigerator(Refrigerator) error {
	fmt.Println("关闭船上冰箱")
	return nil
}
