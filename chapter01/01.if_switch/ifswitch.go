package main

import "fmt"

func main() {
	var rmb int = 21
	var busy bool = false
	switch {
	case rmb <= 20:
		fmt.Println("吃个包子")
	case rmb >= 20 && rmb <= 50:
		fmt.Println("吃顿快餐")
		if busy {
			break
		}
		fmt.Println("再吃顿零食")
	default:
		fmt.Println("我再考虑下")
	}
	fmt.Println("结束")
}
