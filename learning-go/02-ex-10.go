package main

import "fmt"

// Function with variable number of arguments
func printThem(numbers ...int) {
	fmt.Println("About to print numbers")
	for _, d := range numbers {
		fmt.Printf("%d\n", d)
	}
}

func main() {
	printThem(1, 2, 3, 4, 5)
	printThem(6, 7, 8)
}

/* Output:
About to print numbers
1
2
3
4
5
About to print numbers
6
7
8
*/
