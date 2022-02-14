package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := sync.Map{}
	for i := 0; i < 5; i++ {
		go func(i int) {
			m.Store(i, 1)
			for {
				v, _ := m.Load(i)
				//if !ok{
				//	continue
				//}
				m.Store(i, v.(int)+1)
				fmt.Println("i=", i, v)
			}
		}(i)
	}
	time.Sleep(time.Second * 10)
}
