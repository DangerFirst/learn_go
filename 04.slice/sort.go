package main

import (
	"fmt"
)

func main() {
	//定义数组
	sort := []int{3, 5, 1, 2, 4, 7, 7, 3, 6}
	fmt.Println("排序前数组：", sort)
	//排序循环
	for i := 0; i < len(sort)-1; i++ {
		for j := i + 1; j < len(sort); j++ {
			if sort[i] > sort[j] {
				temp := sort[i]
				sort[i] = sort[j]
				sort[j] = temp
			} else {
				continue
			}
		}
	}
	fmt.Println("排序后数组：", sort)
}
