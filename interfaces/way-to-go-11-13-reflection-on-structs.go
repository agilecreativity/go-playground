package main

import (
	"fmt"
	"reflect"
)

type UnknownType struct {
	s1, s2, s3 string
}

func (n UnknownType) String() string {
	return n.s1 + "-" + n.s2 + "-" + n.s3
}

// varialble to investigate
var secret interface{} = UnknownType{"Ada", "Go", "Oberon"}

func main() {
	value := reflect.ValueOf(secret)
	fmt.Printf("value: %v\n", value) // {Ada Go Oberon}

	typ := reflect.TypeOf(secret)
	fmt.Printf("Type: %v\n", typ) // main.UnknownType

	knd := value.Kind()
	fmt.Printf("Kind : %v\n", knd) // struct

	// iterate through the fields of the struct:
	for i := 0; i < value.NumField(); i++ {
		fmt.Println("Field %d: %v\n", i, value.Field(i))
	}

	// Call the first, method which is String()
	results := value.Method(0).Call(nil)
	fmt.Println("Result:", results) // [Ada - Go - Oberon]
}

/* Output:
value: {Ada Go Oberon}
Type: main.UnknownType
Kind : struct
Field %d: %v
 0 Ada
Field %d: %v
 1 Go
Field %d: %v
 2 Oberon
Result: [Ada-Go-Oberon]
*/
