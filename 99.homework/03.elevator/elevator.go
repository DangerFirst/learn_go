package main

type Elevator struct {
	floor        int
	currentFloor int
	task         []int
}

func (e *Elevator) InputToFloor(floor ...int) {
	e.task = append(e.task, floor...)
}

func (e *Elevator) ToFloor() {
	for i, item := range e.task {

	}
}
