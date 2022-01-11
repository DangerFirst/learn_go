package main

import (
	"fmt"
	"sort"
	"time"
)

type Elevator struct {
	floor        int
	currentFloor int
	ElevatorTask
}

type ElevatorTask struct {
	upTask   []int
	downTask []int
}

func (e *Elevator) InputToFloor(floor ...int) {
	task := &e.ElevatorTask
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
	task := []int{}
	if len(e.ElevatorTask.upTask) >= len(e.ElevatorTask.downTask) {
		task = append(e.ElevatorTask.upTask, e.ElevatorTask.downTask...)
	} else {
		task = append(e.ElevatorTask.downTask, e.ElevatorTask.upTask...)
	}
	for _, item := range task {
		time.Sleep(time.Second)
		fmt.Printf("电梯到达第%d层\n", item)
		e.currentFloor = item
	}
}
