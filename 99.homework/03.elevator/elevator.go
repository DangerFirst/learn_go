package main

import (
	"fmt"
	"sort"
	"time"
)

type Elevator struct {
	floor        int
	currentFloor int
	direction    int
	ElevartorDoor
	ElevatorTask
}

type ElevatorTask struct {
	upTask   []int
	downTask []int
}

func (e *Elevator) InputToFloor(floor ...int) {
	task := &e.ElevatorTask
	e.direction = 0
	if floor != nil && floor[0] > e.currentFloor {
		e.direction = 1
	}
	for _, item := range floor {
		if item > e.currentFloor {
			task.upTask = append(task.upTask, item)
		} else {
			task.downTask = append(task.downTask, item)
		}
	}
	sort.Slice(task.upTask, func(i, j int) bool {
		return task.upTask[i] < task.upTask[j]
	})
	sort.Slice(task.downTask, func(i, j int) bool {
		return task.downTask[i] > task.downTask[j]
	})
}

func (e *Elevator) ToFloor() {
	var task []int
	task = append(e.ElevatorTask.downTask, e.ElevatorTask.upTask...)
	if e.direction == 1 {
		task = append(e.ElevatorTask.upTask, e.ElevatorTask.downTask...)
	}
	for _, item := range task {
		time.Sleep(time.Second)
		fmt.Printf("电梯到达第%d层\n", item)
		e.currentFloor = item
		fmt.Print("暂停...")
		e.ElevartorDoor.Switch()
	}
}
