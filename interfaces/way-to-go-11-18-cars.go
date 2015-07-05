// 11.14 Structs, collections and higher order functions
package main

import "fmt"

type Any interface{}

type Car struct {
	Model        string
	Manufacturer string
	BuildYear    int
	// ...
}

type Cars []*Car

// Process all care with the given function f:
func (cs Cars) Process(f func(car *Car)) {
	for _, c := range cs {
		f(c)
	}
}

// Find all cars matching a given criteria
func (cs Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make([]*Car, 0)
	cs.Process(func(c *Car) {
		if f(c) {
			cars = append(cars, c)
		}
	})
	return cars
}

// Process cars and create new data
func (cs Cars) Map(f func(car *Car) Any) []Any {
	result := make([]Any, 0)
	ix := 0
	cs.Process(func(c *Car) {
		result[ix] = f(c)
		ix++
	})
	return result
}

func MakeSortedAppender(manufactures []string) (func(car *Car), map[string]Cars) {
	// Prepare maps of sorted cars
	sortedCars := make(map[string]Cars)

	for _, m := range manufactures {
		sortedCars[m] = make([]*Car, 0)
	}

	sortedCars["Default"] = make([]*Car, 0)

	// Prepare appender function
	appender := func(c *Car) {
		if _, ok := sortedCars[c.Manufacturer]; ok {
			sortedCars[c.Manufacturer] = append(sortedCars[c.Manufacturer], c)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], c)
		}
	}

	return appender, sortedCars
}

func main() {
	// make some cars
	ford := &Car{"Fiesta", "Ford", 2008}
	bmw := &Car{"XL 450", "BMW", 2011}
	merc := &Car{"D600", "Mercedes", 2009}
	bmw2 := &Car{"X 800", "BMW", 2008}

	allCars := Cars([]*Car{ford, bmw, merc, bmw2})

	allNewBMW := allCars.FindAll(func(car *Car) bool {
		return (car.Manufacturer == "BMW") && (car.BuildYear > 2010)
	})

	fmt.Printf("All Cars: %\n", allCars)
	fmt.Println("\n----------------------")
	fmt.Printf("New BMWs: %\n", allNewBMW)

	manufacturers := []string{
		"Ford",
		"Aston Martin",
		"Land Rover",
		"BMW",
		"Jaguar",
	}
	sortedAppender, sortedCars := MakeSortedAppender(manufacturers)
	allCars.Process(sortedAppender)

	fmt.Println("\n Map sortedCars: ", sortedCars)
	BMWCount := len(sortedCars["BMW"])
	fmt.Println("We have", BMWCount, "BMWs")
}

/* Output:
All Cars: [%!
(*main.Car=&{Fiesta Ford 2008}) %!
(*main.Car=&{XL 450 BMW 2011}) %!
(*main.Car=&{D600 Mercedes 2009}) %!
(*main.Car=&{X 800 BMW 2008})]
----------------------
New BMWs: [%!
(*main.Car=&{XL 450 BMW 2011})]
 Map sortedCars:  map[BMW:[0x2081fa420 0x2081fa480] Jaguar:[] Default:[0x2081fa450] Ford:[0x2081fa3f0] Aston Martin:[] Land Rover:[]]
We have 2 BMWs
*/
