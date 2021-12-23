package main

import (
	"fmt"
	_ "github.com/spf13/cobra"
	"os"
)

func main() {
	var (
		name   string
		sex    string
		age    string //int
		tall   string //float64
		weight string //float64
	)
	arguments := os.Args
	name = arguments[1]
	sex = arguments[2]
	age = arguments[3]
	tall = arguments[4]
	weight = arguments[5]
	fmt.Println(arguments, len(arguments), cap(arguments))
	fmt.Println("name：", name)
	fmt.Println("sex：", sex)
	fmt.Println("age：", age)
	fmt.Println("tall：", tall)
	fmt.Println("weight：", weight)
}
