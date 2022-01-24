package main

import "fmt"

func main() {
	security := Assets{assets: []Asset{
		&GlassDoor{},
		&WoodDoor{},
	}}
	fmt.Println("开始上班")
	security.Unlock()
	security.DoStarWork()
	fmt.Println("8小时工作，准备下班")
	security.DoStopWork()
	security.Lock()
	fmt.Println("Done")
}
