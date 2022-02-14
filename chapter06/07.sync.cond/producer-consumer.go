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
	PCond     *sync.Cond
	CCond     *sync.Cond
}

type Producer struct {
}

func (Producer) Produce(s *Store) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	if s.DataCount == s.Max {
		fmt.Println("生产者在等消费")
		s.PCond.Wait()
	}
	fmt.Println("生产者开始生产+1")
	s.DataCount++
	s.CCond.Signal()
}

type Consumer struct {
}

func (Consumer) Consume(s *Store) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	if s.DataCount == 0 {
		fmt.Println("消费者在等库存")
		s.CCond.Wait()
	}
	fmt.Println("消费者开始消费-1")
	s.DataCount--
	s.PCond.Signal()
}

func main() {
	s := &Store{
		Max: 10,
	}
	s.PCond = sync.NewCond(&s.Lock)
	s.CCond = sync.NewCond(&s.Lock)
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
