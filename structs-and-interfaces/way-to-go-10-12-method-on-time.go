package main

import (
	"fmt"
	"time"
)

type myTime struct {
	time.Time // anonymous field
}

func (t myTime) first3Chars() string {
	return t.Time.String()[0:3]
}

func main() {
	m := myTime{time.Now()}

	// calling existing String method on anonymous Time field
	fmt.Println("Full time now:", m.String())

	// calling myTime.m.first3Chars()
	fmt.Println("First 3 chars:", m.first3Chars())
}

/* Output:
Full time now: 2015-07-06 22:51:04.243377957 +1000 AEST
First 3 chars: 201
*/
