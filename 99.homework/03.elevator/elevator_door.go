package main

import "fmt"

type ElevatorDoor struct {
}

func (d *ElevatorDoor) Switch() {
	fmt.Print("开门...")
	fmt.Println("关门...")
}
