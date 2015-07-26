package main

import "fmt"

type Any interface{}

type stack []Any

func (s stack) Empty() bool {
	return len(s) == 0
}

func (s stack) Peek() Any {
	return s[len(s)-1]
}

func (s *stack) Put(i Any) {
	(*s) = append((*s), i)
}

func (s *stack) Pop() Any {
	d := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return d
}

func main() {
	var s stack
	for i := 0; i < 3; i++ {
		s.Put(i)
		fmt.Printf("len=%d\n", len(s))
		fmt.Printf("peek=%d\n", s.Peek())
	}
}
