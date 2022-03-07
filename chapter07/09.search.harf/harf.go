package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

var totalCompare int = 0

func main() {
	//size:=1000
	//arr := generateRandomData(size)
	//arr := sampleData
	//fmt.Println(arr)

	newArr := append([]int64{}, sampleData...)

	startTime := time.Now()

	quickSort(&newArr, 0, len(newArr)-1)
	fmt.Println(newArr)
	for i := 0; i < 10000; i++ {
		fmt.Println(search(&newArr, 11000))
		fmt.Println(search(&newArr, 3016))
		fmt.Println(search(&newArr, 2))
		fmt.Println(search(&newArr, 44))
	}
	finishTime := time.Now()
	fmt.Println("总次数：", totalCompare)
	fmt.Println("总用时：", finishTime.Sub(startTime))
}

func search(arrP *[]int64, targetNum int64) bool {
	return searchHarf(arrP, 0, len(*arrP)-1, targetNum)
}

func searchHarf(arrP *[]int64, left, right int, targetNum int64) bool {
	middleIndex := (left + right) / 2
	data := (*arrP)[middleIndex]
	totalCompare++
	if data < targetNum {
		if left == right {
			return false
		}
		return searchHarf(arrP, middleIndex+1, right, targetNum)
	} else if data > targetNum {
		if left == right || right == 1 {
			return false
		}
		return searchHarf(arrP, left, middleIndex-1, targetNum)
	} else {
		return true
	}
}

func generateRandomData(size int) []int64 {
	arr := make([]int64, 0, size)

	for i := 0; i < size; i++ {
		i, _ := rand.Int(rand.Reader, big.NewInt(10000))
		arr = append(arr, i.Int64())
	}
	return arr
}
