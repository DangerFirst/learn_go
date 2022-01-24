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

func (r Runner)Run()  {
	fmt.Println(r.name,"开始跑")
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Int()%10)*time.Second)
	fmt.Println(r.name,"跑到终点")
}

func main() {
	runnerCount:=10
	runner:=[]Runner{}

	wg:=sync.WaitGroup{}
	wg.Add(runnerCount)

	for i:=0;i<runnerCount;i++{

	}

	wg.Wait()
}