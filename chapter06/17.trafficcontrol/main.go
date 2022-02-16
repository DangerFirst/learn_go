package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 200; i++ {
		go sendRequest()
	}
	time.Sleep(5 * time.Second)
}

func sendRequest() {
	for {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Int63()%30) * time.Millisecond)
		serviceRequest()
	}
}

var jobControlCh = make(chan struct{}, 10)

func serviceRequest() {
	accept := trafficControlStart()
	if accept {
		defer trafficControlFinish()
		fmt.Println("服务请求")
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Int63()%30) * time.Millisecond)
	} else {
		fmt.Println("服务请求被拒绝")
	}
}

func trafficControlStart() (accept bool) {
	select {
	case jobControlCh <- struct{}{}:
		fmt.Println("接受请求")
		return true
	default:
		fmt.Println("拒绝请求")
		return
	}
}

func trafficControlFinish() {
	<-jobControlCh
}
