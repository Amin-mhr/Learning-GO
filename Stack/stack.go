package main

import "fmt"

type Stack struct {
	data []interface{}
}

func (s *Stack) Pop() interface{} {
	if len(s.data) == 0 {
		return nil
	} else {
		v := s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]
		return v
	}
}

func (s *Stack) Push(v interface{}) {
	s.data = append(s.data, v)
}

func (s *Stack) Length() int {
	return len(s.data)
}

func main() {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	fmt.Println(stack.Length())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Length())
	fmt.Println(stack.Pop())
}
