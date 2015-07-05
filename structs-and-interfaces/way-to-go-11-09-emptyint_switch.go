package main

import (
	"fmt"
)

type specialString string

func TypeSwitch(testValue specialString) {
	testFunc := func(any interface{}) {
		switch v := any.(type) {
		case bool:
			fmt.Printf("any %v is a bool type", v)
		case int:
			fmt.Printf("any %v is an int type", v)
		case float32:
			fmt.Printf("any %v is a float32 type ", v)
		case specialString:
			fmt.Printf("any %v is a special string!", v)
		default:
			fmt.Println("Unknown type %v\n", v)
		}
	}
	testFunc(testValue)
}

func main() {
	var whatIsThis specialString = "Hello"
	TypeSwitch(whatIsThis)
}

// Output:
// any Hello is a special string!
