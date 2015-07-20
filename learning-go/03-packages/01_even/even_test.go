package even

import (
	"fmt"
	"testing"
)

func TestEven(t *testing.T) {
	if !Even(2) {
		t.Log("2 should be even!")
		t.Fail()
	}
}

func ExampleEven() {
	if Even(2) {
		fmt.Printf("Number %d is even\n", 2)
	}
	// Output:
	// Number 2 is even
}
