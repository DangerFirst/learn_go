package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		//countDict()
		//countDictGoroutineSafe()
		countDictGoroutineSafeRW()
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

func countDictGoroutineSafeRW() {
	fmt.Println("开始数")
	TotalCount := 0
	TotalCountLock := sync.RWMutex{}
	wg := sync.WaitGroup{}
	wg.Add(10)
	for p := 0; p < 5; p++ {
		go func(p int) {
			defer wg.Done()
			fmt.Println(p, "读锁开始时间：", time.Now())
			TotalCountLock.RLock()
			fmt.Println(p, "读锁拿到锁时间：", time.Now())
			time.Sleep(time.Second)
			defer TotalCountLock.RUnlock()
			//fmt.Print("正在统计第", p, "页，")
			TotalCount += 100
		}(p)
	}

	for p := 0; p < 5; p++ {
		go func(p int) {
			defer wg.Done()
			fmt.Println(p, "写锁开始时间：", time.Now())
			TotalCountLock.Lock()
			fmt.Println(p, "写锁拿到锁时间：", time.Now())
			defer TotalCountLock.Unlock()
			//fmt.Print("正在统计第", p, "页，")
			TotalCount += 100
		}(p)
	}
	wg.Wait()

}
