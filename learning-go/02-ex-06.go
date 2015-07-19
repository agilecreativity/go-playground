package main

import (
	"fmt"
	"strconv"
)

// Use of named return parameter
func average(xs []float64) (avg float64) {
	sum := 0.0
	switch len(xs) {
	case 0:
		avg = 0
	default:
		for _, v := range xs {
			sum += v
		}
		avg = sum / float64(len(xs))
	}
	return
}

// Return multiple values
func order(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

// Note: stack is not exported
type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(k int) {
	if s.i+1 > 9 {
		return
	}
	s.data[s.i] = k
	s.i++
}

func (s *stack) pop() int {
	s.i--
	return s.data[s.i]
}

func (s stack) String() string {
	var str string
	for i := 0; i < s.i; i++ {
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}

func main() {
	var s stack
	s.push(25)
	fmt.Printf("Stack: %v\n", s)
	s.push(24)
	fmt.Printf("Stack: %v\n", s)
}

/* Output:
Stack: [0:25]
Stack: [0:25][1:24]
*/
