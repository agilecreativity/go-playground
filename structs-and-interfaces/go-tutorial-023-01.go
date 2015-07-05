package main

import "fmt"

// empty interface
func whatIsThis(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("It is a string: %s\n", v)
	case uint32:
		fmt.Printf("It is %T with value: %d\n", v, v)
	default:
		fmt.Printf("Type is %T, with value: %v\n", v, v)
	}
}

func main() {
	whatIsThis([...]string{"The", "quick", "brown", "fox", "..."})
	whatIsThis("Hello")
	whatIsThis(uint32(42.0))
	whatIsThis(42.356)
}

/* Output:
Type is [5]string, with value: [The quick brown fox ...]
It is a string: Hello
It is uint32 with value: 42
Type is float64, with value: 42.356
*/
