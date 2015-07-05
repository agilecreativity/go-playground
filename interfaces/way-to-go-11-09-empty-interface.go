package main

import (
	"fmt"
)

var i = 5
var str = "ABC"

type Person struct {
	name string
	age  int
}

type Any interface {
}

func printType(val Any) {
	switch t := val.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case bool:
		fmt.Printf("Type boolean %T\n", t)
	case Person:
		fmt.Printf("Type Person %T\n", t)
	case *Person:
		fmt.Printf("Type pointer to Person %T\n", t)
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
}

func main() {
	var val Any
	val = 5
	fmt.Printf("val has the value: %v\n", val)

	val = str
	fmt.Printf("val has the value: %v\n", val)

	person := Person{
		name: "rob pike",
		age:  55,
	}

	val = person
	fmt.Printf("val has the value: %v\n", val)

	printType(val)

	geek := new(Person)
	geek.name = "Ken Thompson"
	geek.age = 30
	val = geek
	printType(val)
}

// Output:
// val has the value: 5
// val has the value: ABC
// val has the value: {rob pike 55}
// Type Person main.Person
// Type pointer to Person *main.Person
