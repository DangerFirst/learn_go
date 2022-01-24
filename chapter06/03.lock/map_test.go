package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	m := map[int]int{}
	mLock := sync.Mutex{}
	for i := 0; i < 5; i++ {
		go func() {
			for {
				mLock.Lock()
				v := m[i]
				m[i] = v + 1
				fmt.Println("i=", i, m[i])
				mLock.Unlock()
			}
		}()
	}
	time.Sleep(time.Second * 1)
}
