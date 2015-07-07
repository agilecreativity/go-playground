package main

import (
	"fmt"
)

type Camera struct{}

func (c *Camera) TakeAPicture() string {
	return "Click"
}

type Phone struct{}

func (p *Phone) Call() string {
	return "Ring Ring"
}

// multiple inheritance
type CameraPhone struct {
	Camera
	Phone
}

func main() {
	cp := new(CameraPhone)
	fmt.Println("Our new CameraPhone exhibits multiple behaviors...")
	fmt.Println("It exhibits behavior of a Camera:", cp.TakeAPicture())
	fmt.Println("It works like a phone too:", cp.Call())
}

/* Output:
Our new CameraPhone exhibits multiple behaviors...
It exhibits behavior of a Camera: Click
It works like a phone too: Ring Ring
*/
