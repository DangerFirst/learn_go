package main

import "fmt"

func main() {
	elevator := &Elevator{
		floor:        5,
		currentFloor: 3,
	}
	fmt.Println("电梯当前楼层为：", elevator.currentFloor)
	elevator.InputToFloor(5, 1, 3, 4)
	elevator.ToFloor()
	fmt.Printf("电梯停在第%d层\n", elevator.currentFloor)
}
