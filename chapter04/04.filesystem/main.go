package main

import (
	"fmt"
)

func main() {
	var data string
	{
		var equipment IOInterface = &Paper{}
		data = equipment.read()
		fmt.Println(data)
	}
	{
		var equipment IOInterface = &Soft{}
		data = equipment.read()
		fmt.Println(data)
	}
}

type IOInterface interface {
	read() string
}

type Paper struct {
}

func (Paper) read() string {
	return "纸带数据"
}

type Soft struct {
}

func (Soft) read() string {
	return "软件数据"
}
