package main

import (
	"fmt"
	"testing"
	"time"
)

func TestDefChannel(t *testing.T) {
	intCh := make(chan int, 1)
	fmt.Println("intCh:", intCh)
	fmt.Println("装入数据")
	intCh <- 3
	fmt.Println("intCh:", intCh)
	fmt.Println("取出数据")
	outCh := <-intCh
	fmt.Println("outCh:", outCh)
}

func TestChanPutGet(t *testing.T) {
	intCh := make(chan int) //创建一个不带size的channel（不带buffer）
	workerCount := 10
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			fmt.Println("开始存入@", time.Now())
			intCh <- i
			fmt.Println("结束存入@", time.Now())
		}(i)
	}
	fmt.Println(time.Now())
	time.Sleep(2 * time.Second)
	fmt.Println(time.Now())
	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh
			fmt.Printf("出%d,拿到%d\n", o, out)
		}(o)
	}
	time.Sleep(1 * time.Second)
	fmt.Println(time.Now())
}

func TestChanPutGet2(t *testing.T) {
	intCh := make(chan int, 10) //创建一个带size的channel（带buffer）
	workerCount := 10
	for i := 0; i < workerCount; i++ {
		go func(i int) {
			fmt.Println("开始存入@", time.Now())
			intCh <- i
			fmt.Println("结束存入@", time.Now())
		}(i)
	}
	fmt.Println(time.Now())
	time.Sleep(2 * time.Second)
	fmt.Println(time.Now())
	for o := 0; o < workerCount; o++ {
		go func(o int) {
			out := <-intCh
			fmt.Printf("出%d,拿到%d\n", o, out)
		}(o)
	}
	time.Sleep(1 * time.Second)
	fmt.Println(time.Now())
}

func TestRangeChannel(t *testing.T) {
	intChan := make(chan int, 10)
	intChan <- 1
	intChan <- 2
	intChan <- 3
	intChan <- 4
	intChan <- 5
	for o := range intChan {
		fmt.Println(o)
	}
}

func TestRangeClosedChannel(t *testing.T) {
	intChan := make(chan int, 10)
	intChan <- 1
	intChan <- 2
	intChan <- 3
	intChan <- 4
	intChan <- 5
	close(intChan)
	o1, ok := <-intChan
	fmt.Println(o1, ok)
	for o := range intChan {
		fmt.Println(o)
	}
	o2, ok := <-intChan
	fmt.Println(o2, ok)
}

func TestSelectChannel(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		//time.Sleep(1*time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "go"
	}()

	select {
	case o := <-ch1:
		fmt.Println("ch1 ready,go", o)
	case o := <-ch2:
		fmt.Println("ch2 ready,go", o)
	default:
		fmt.Println("所有的channel未准备好，走default")
	}
	fmt.Println("DONE")
}

func TestMultipleSelect(t *testing.T) {
	ch1 := make(chan int)
	for i := 0; i < 10; i++ {
		go func(i int) {
			select {
			case <-ch1:
				fmt.Println(time.Now(), i)
			}
		}(i)
	}
	fmt.Println("关闭channel", time.Now())
	close(ch1)
	time.Sleep(1 * time.Second)
}
