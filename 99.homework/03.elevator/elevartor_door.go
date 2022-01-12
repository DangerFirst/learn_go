package main

import "fmt"

type ElevartorDoor struct {
}

func (d *ElevartorDoor) Switch() {
	fmt.Print("开门...")
	fmt.Println("关门...")
}
