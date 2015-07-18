// Generic map function in Go
package main

import "fmt"

// define empty interface as a type
type e interface{}

func doubleUp(f e) e {
	switch f.(type) {
	case int:
		return f.(int) * 2
	case string:
		return f.(string) + f.(string)
	}
	return f
}

func Map(n []e, f func(e) e) []e {
	m := make([]e, len(n))
	for k, v := range n {
		m[k] = f(v)
	}
	return m
}

func main() {
	m := []e{1, 2, 3, 4}
	s := []e{"a", "b", "c", "d"}

	mf := Map(m, doubleUp)
	fmt.Println(mf)
	sf := Map(s, doubleUp)

	fmt.Printf("%v\n", mf)
	fmt.Printf("%v\n", sf)
}

/* Output:
[2 4 6 8]
[2 4 6 8]
[aa bb cc dd]
*/