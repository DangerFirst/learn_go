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
	startTime := time.Now()
	arr := sampleData
	fmt.Println(arr)
	for i := 0; i < 10000; i++ {
		fmt.Println(search(&arr, 110))
		fmt.Println(search(&arr, 3016))
		fmt.Println(search(&arr, 2))
		fmt.Println(search(&arr, 44))
	}
	finishTime := time.Now()
	fmt.Println("总次数：", totalCompare)
	fmt.Println("总用时：", finishTime.Sub(startTime))
}

func search(arrP *[]int64, targetNum int64) bool {
	for _, v := range *arrP {
		totalCompare++
		if v == targetNum {
			return true
		}
	}
	return false
}

func generateRandomData(size int) []int64 {
	arr := make([]int64, 0, size)

	for i := 0; i < size; i++ {
		i, _ := rand.Int(rand.Reader, big.NewInt(10000))
		arr = append(arr, i.Int64())
	}
	return arr
}
