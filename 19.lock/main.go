package main

import (
	"fmt"
	"sync"
)

func main() {
	for i := 0; i < 5; i++ {
		//countDict()
		countDictGoroutineSafe()
	}
}

func countDict() {
	fmt.Println("开始数")
	TotalCount := 0
	wg := sync.WaitGroup{}
	wg.Add(5000)
	for p := 0; p < 5000; p++ {
		go func() {
			defer wg.Done()
			//fmt.Print("正在统计第", p, "页，")
			TotalCount += 100
		}()
	}
	wg.Wait()
	fmt.Println("预计有", 5000*100, "个字")
	fmt.Println("总共有", TotalCount, "个字")
}

func countDictGoroutineSafe() {
	fmt.Println("开始数")
	TotalCount := 0
	TotalCountLock := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(5000)
	for p := 0; p < 5000; p++ {
		go func() {
			defer wg.Done()
			TotalCountLock.Lock()
			defer TotalCountLock.Unlock()
			//fmt.Print("正在统计第", p, "页，")
			TotalCount += 100
		}()
	}
	wg.Wait()
	fmt.Println("预计有", 5000*100, "个字")
	fmt.Println("总共有", TotalCount, "个字")
}
