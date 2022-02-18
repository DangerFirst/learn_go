package main

import (
	"fmt"
	"math/rand"
)

func bubble(arr *[]int) {
	for i := 0; i < len(*arr)-1; i++ {
		for j := i + 1; j < len(*arr); j++ {
			if (*arr)[i] > (*arr)[j] {
				(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
			}
		}
		fmt.Println("中间状态:", *arr)
	}
	fmt.Println(*arr)
}

func main() {
	arrSize := 10
	arr := []int{}
	for i := 0; i < arrSize; i++ {
		arr = append(arr, rand.Intn(50))
	}
	fmt.Println("排序前：", arr)
	bubble(&arr)
}
