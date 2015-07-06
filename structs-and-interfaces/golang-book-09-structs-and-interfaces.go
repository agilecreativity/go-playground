// See: https://www.golang-book.com/books/intro/9
package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

// embedded type
type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Printf("Hi, my name is %s\n", p.Name)
}

type Android struct {
	Person
	Model string
}

type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

// Use interface as fields:
type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Printf("Circle Area: %f\n", c.area())

	r := Rectangle{x1: 0, y1: 0, x2: 10, y2: 10}
	fmt.Printf("Rectangle area: %f\n", r.area())

	// Use interface
	fmt.Printf("Total area: %f\n", totalArea(&c, &r))

	// Using MultiShape
	shapes := &MultiShape{[]Shape{&c, &r}}

	fmt.Printf("MultiShape area: %f\n", shapes.area())

	// using embedded struct
	a := new(Android)
	a.Person.Name = "R2D2"
	a.Person.Talk()
}
