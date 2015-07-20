// unction that return functions
package main

import "fmt"

// Function that return function
func plusTwo() func(int) int {
	return func(x int) int {
		return x + 2
	}
}

// Generic functions that return function
func plusX(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	p2 := plusTwo()
	fmt.Printf("Pluse Two: %d\n", p2(2))

	plusTen := plusX(10)
	fmt.Printf("Plus Ten: %d\n", plusTen(3))
}

/* Output:
Pluse Two: 4
Plus Ten: 13
*/
