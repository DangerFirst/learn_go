package main

import "fmt"

var _ Door = &WoodDoor{}

type WoodDoor struct{}

func (*WoodDoor) Unlock() {
	fmt.Println("WoodDoor Unlock")
}

func (*WoodDoor) Lock() {
	fmt.Println("WoodDoor Lock")
}

func (*WoodDoor) Open() {
	fmt.Println("WoodDoor Open")
}

func (*WoodDoor) Close() {
	fmt.Println("WoodDoor Close")
}
