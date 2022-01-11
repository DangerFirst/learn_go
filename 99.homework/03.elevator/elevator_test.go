package main

import "testing"

func TestCase1(t *testing.T) {
	e := &Elevator{
		floor:        5,
		currentFloor: 1,
	}
	e.InputToFloor()
	e.ToFloor()
	if e.currentFloor != 1 {
		t.Fatalf("电梯应该停留在1层，实际停留在%d层", e.currentFloor)
	}
}

func TestCase2(t *testing.T) {
	e := &Elevator{
		floor:        5,
		currentFloor: 1,
	}
	e.InputToFloor(3)
	e.ToFloor()
	if e.currentFloor != 3 {
		t.Fatalf("电梯应该停留在3层，实际停留在%d层", e.currentFloor)
	}
}

func TestCase3(t *testing.T) {
	e := &Elevator{
		floor:        5,
		currentFloor: 3,
	}
	e.InputToFloor(4, 2)
	e.ToFloor()
	if e.currentFloor != 2 {
		t.Fatalf("电梯应该停留在2层，实际停留在%d层", e.currentFloor)
	}
}
