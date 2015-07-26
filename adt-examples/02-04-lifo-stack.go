// Inspired by:
// https://gist.github.com/bemasher/1777766
// https://github.com/alinz/gonote/blob/master/note/util/stack.go
package main

import (
	"fmt"
)

type Any interface{}

type Element struct {
	value Any
	next  *Element
}

type Stack struct {
	top  *Element
	size int
}

// Return the stack's length
func (s *Stack) Len() int {
	return s.size
}

// Push a new element onto the stack
func (s *Stack) Push(value Any) {
	s.top = &Element{value, s.top}
	s.size++
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *Stack) Pop() (value Any) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

// Peeks returns a top without removing it from the list
func (s *Stack) Peek() (value Any, exists bool) {
	exists = false

	if s.size > 0 {
		value = s.top.value
		exists = true
	}

	return
}

func main() {
	stack := new(Stack)

	stack.Push("One")
	stack.Push("Two")
	value, found := stack.Peek()

	if found {
		fmt.Printf("Peek value: %s\n", value)
	} else {
		fmt.Println("No value found!")
	}

	stack.Push("Three")
	stack.Push("Four")

	for stack.Len() > 0 {
		fmt.Printf("%s\n", stack.Pop())
	}
	fmt.Println()
}

/* Output:
Peek value: Two
Four
Three
Two
One
*/
