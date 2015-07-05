package main

import "fmt"

// empty interface
func whatIsThis(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("It is a string: %s\n", i.(string))
	case uint32:
		fmt.Printf("It is an unsigned 32-bit integer: %d\n", i.(uint32))
	default:
		fmt.Printf("Something else")
	}
}

func main() {
	whatIsThis([...]string{"The", "quick", "brown", "fox", "..."})
	whatIsThis("Hello")
	whatIsThis(uint32(42.0))
	whatIsThis(42.356)
}

/* Output:
Something elseIt is a string: Hello
It is an unsigned 32-bit integer: 42
Something else
*/
