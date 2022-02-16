package main

import (
	"fmt"
	"sync"
)

type Stack struct {
	sync.Mutex
	data []interface{}
}

func (s *Stack) Push(date interface{}) {
	s.Lock()
	defer s.Unlock()
	s.data = append([]interface{}{date}, s.data...)
}

func (s *Stack) Pop() (interface{}, bool) {
	s.Lock()
	defer s.Unlock()
	if len(s.data) > 0 {
		o := s.data[0]
		s.data = s.data[1:]
		return o, true
	}
	return nil, false
}

func main() {
	//Add(1,2)
	s := &Stack{}
	s.Push(111)
	s.Push(222)
	s.Push(333)
	s.Push(nil)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

//func Add(a, b int) {
//	//Add(a,b)
//}
