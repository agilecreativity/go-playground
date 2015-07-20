package main

import even "./01_even"
import "fmt"

func main() {
	i, j := 5, 6

	fmt.Printf("Is %d even? %v\n", i, even.Even(i))
	fmt.Printf("Is %d odd?  %v\n", i, even.Odd(i))
	fmt.Printf("Is %d even?  %v\n", j, even.Even(j))
	fmt.Printf("Is %d odd?  %v\n", j, even.Odd(j))
}

/* Output:
Is 5 even? false
Is 5 odd?  true
Is 6 even?  true
Is 6 odd?  false
*/
