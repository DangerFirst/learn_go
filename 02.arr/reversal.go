package main

import "fmt"

func main() {
	array := [3]int{1, 2, 3}
	arrayRe := [3]int{}
	fmt.Println("反转前数组为：", array)
	for i, val := range array {
		arrayRe[len(array)-i-1] = val
	}
	fmt.Println("反转后数组为", arrayRe)
}
