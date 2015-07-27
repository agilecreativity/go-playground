// See: https://github.com/hishboy/gocommons/blob/master/lang/stack.go
package main

import (
	"fmt"
	"sync"
)

type Any interface{}

type Node struct {
	data Any
	next *Node
}

type Stack struct {
	head  *Node
	count int
	lock  *sync.Mutex
}

func NewStack() *Stack {
	s := &Stack{}
	s.lock = &sync.Mutex{}
	return s
}

func (s *Stack) Len() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.count
}

func (s *Stack) Push(item Any) {
	s.lock.Lock()
	defer s.lock.Unlock()

	n := &Node{data: item}

	if s.head == nil {
		s.head = n
	} else {
		n.next = s.head
		s.head = n
	}
	s.count++
}

func (s *Stack) Pop() Any {
	s.lock.Lock()
	defer s.lock.Unlock()

	var n *Node

	if s.head != nil {
		n = s.head
		s.head = n.next
		s.count--
	}

	if n == nil {
		return nil
	}

	return n.data
}

func (s *Stack) Peek() Any {
	s.lock.Lock()
	defer s.lock.Unlock()

	n := s.head

	if n == nil || n.data == nil {
		return nil
	}
	return n.data
}

func main() {
	s := NewStack()
	s.Push("1st")
	s.Push("2nd")
	s.Push("3rd")
	s.Push("4th")
	fmt.Printf("Peek             : %s\n", s.Peek())
	fmt.Printf("Len (before pop) : %d\n", s.Len())
	fmt.Printf("Pop item         : %s\n", s.Pop())
	fmt.Printf("Len (after pop)  : %d\n", s.Len())
	fmt.Printf("Peek             : %s\n", s.Peek())
}

/* Output:
Peek             : 4th
Len (before pop) : 4
Pop item         : 4th
Len (after pop)  : 3
Peek             : 3rd
*/
