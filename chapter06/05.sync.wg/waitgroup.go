package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Runner struct {
	name string
}

func (r Runner) Run(startPointWg, wg *sync.WaitGroup) {
	defer wg.Done()
	startPointWg.Wait()
	start := time.Now()
	fmt.Println(r.name, "开始跑@", start)
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Int()%10) * time.Second)
	finsh := time.Now()
	fmt.Println(r.name, "跑到终点,用时：", finsh.Sub(start))
}

func main() {
	runnerCount := 10
	runners := []Runner{}

	wg := sync.WaitGroup{}
	wg.Add(runnerCount)

	startPointWg := sync.WaitGroup{}
	startPointWg.Add(1)

	for i := 0; i < runnerCount; i++ {
		runners = append(runners, Runner{
			name: fmt.Sprint(i),
		})
	}

	for _, runnerItem := range runners {
		go runnerItem.Run(&startPointWg, &wg)
	}

	startPointWg.Done()
	wg.Wait()
	fmt.Println("赛跑结束")
}
