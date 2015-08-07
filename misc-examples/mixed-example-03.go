package main

import (
	"fmt"
	"sort"
)

type Animal interface {
	Name() string
}

type Dog struct{}

func (dog Dog) Name() string {
	return "Jazzi"
}

func main() {
	var a Animal
	a = &Dog{}
	fmt.Println(a.Name()) //=> "Jazzi"

	// sorting
	strs := []string{
		"the",
		"quick",
		"brown",
		"fox",
	}

	fmt.Printf("Text before sort: %v\n", strs)
	sort.Strings(strs)
	fmt.Printf("Text after sort : %v\n", strs)
}

/* Output:
Jazzi
Text before sort: [the quick brown fox]
Text after sort : [brown fox quick the]
*/
