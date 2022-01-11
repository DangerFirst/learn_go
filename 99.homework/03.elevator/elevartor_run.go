package main

import "fmt"

func main() {
	elevator := &Elevator{
		floor:        5,
		currentFloor: 3,
	}
	fmt.Println("电梯当前楼层为：", elevator.currentFloor)
	elevator.InputToFloor(4, 1, 2, 5)
	elevator.ToFloor()
	fmt.Printf("电梯停在：%d层\n", elevator.currentFloor)
}
