package main

import "fmt"

func main() {
	plus := func(a, b int) int {
		return a + b
	}

	fmt.Println(plus(2, 3)) //=> 5

	// Pointer examples
	i := 1
	var ip *int = &i
	fmt.Println("Value  :", i)
	fmt.Println("Pointer:", &i)
	fmt.Println("Pointer:", ip)

	// Structs
	type Person struct {
		Name string
		Age  int
	}

	p1 := Person{"John Carmack", 50}
	p2 := Person{Name: "John Carmack", Age: 50}
	p3 := &Person{"John Carmack", 50}
	p4 := &Person{Age: 50}

	fmt.Printf("Person 1 :%v\n", p1)
	fmt.Printf("Person 2 :%v\n", p2)
	fmt.Printf("Person 3 :%v\n", p3)
	fmt.Printf("Person 4 :%v\n", p4)

	// new(T) : returns a pointer to a newly allocated zero value of type T (e.g. *T).
	// make(T, len, cap) : creates slices, maps, and channels only, and
	// it returns an initialized (not zeroed) value of type T (not *T)

	p5 := new(Person) // *Person
	p5.Name = "John Smith"
	fmt.Printf("Person 5 :%v\n", p5) //Note: the value for p.Age is set to 0 by default

	nums := make([]int, 10) // slice of []int length 10
	fmt.Printf("slice of []int:%v\n", nums)
}

/* Output:
5
Value  : 1
Pointer: 0x82024e240
Pointer: 0x82024e240
Person 1 :{John Carmack 50}
Person 2 :{John Carmack 50}
Person 3 :&{John Carmack 50}
Person 4 :&{ 50}
Person 5 :&{John Smith 0}
slice of []int:[0 0 0 0 0 0 0 0 0 0]
*/
