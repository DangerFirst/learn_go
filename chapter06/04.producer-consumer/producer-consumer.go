package main

import (
	"fmt"
	"sync"
	"time"
)

type Store struct {
	DataCount int
	Max       int
	Lock      sync.Mutex
}

type Producer struct {
}

func (Producer) Produce(s *Store) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	if s.DataCount == s.Max {
		fmt.Println("生产者看到库存满了，不生产")
		return
	}
	fmt.Println("生产者开始生产+1")
	s.DataCount++
}

type Consumer struct {
}

func (Consumer) Consume(s *Store) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	if s.DataCount == 0 {
		fmt.Println("消费者看到库存没了，不消费")
		return
	}
	fmt.Println("消费者开始消费-1")
	s.DataCount--
}

func main() {
	s := &Store{
		Max: 10,
	}
	pCount, cCount := 50, 50
	for i := 0; i < pCount; i++ {
		go func() {
			for {
				time.Sleep(100 * time.Millisecond)
				Producer{}.Produce(s)
			}
		}()
	}
	for i := 0; i < cCount; i++ {
		go func() {
			for {
				time.Sleep(100 * time.Millisecond)
				Consumer{}.Consume(s)
			}
		}()
	}
	time.Sleep(time.Second)
}
