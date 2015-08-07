package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{2, 3}
	fmt.Printf("Area of rectangle of side 2 by 3 is  :%d\n", r.area())
	fmt.Printf("Perim of rectangle of side 2 by 3 is :%d\n", r.perim())
}
