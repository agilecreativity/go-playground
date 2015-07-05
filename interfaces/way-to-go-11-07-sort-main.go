package main

import "fmt"
import sort "./s11_07_sort"

func ints() {
	data := []int{74, 59, 238, -784, 9689, 959, 905, 0, 0, 42, 1234}
	a := sort.IntArray(data)
	sort.Sort(a)

	if sort.IsSorted(a) {
		fmt.Printf("The sorted array is : %v\n", a)
	} else {
		panic("Problem sorting the array")
	}
}

func strings() {
	data := []string{
		"monday", "tuesday", "wednesday", "thurday", "friday", "saturday", "sunday", "the", "quick", "brown", "fox",
	}

	a := sort.StringArray(data)
	sort.Sort(a)
	if !sort.IsSorted(a) {
		panic("IsSorted() is not working correctly!")
	}

	fmt.Printf("The sorted array is: %v\n", a)
}

type day struct {
	num       int
	shortName string
	longName  string
}

type dayArray struct {
	data []*day
}

func (p *dayArray) Len() int           { return len(p.data) }
func (p *dayArray) Less(i, j int) bool { return p.data[i].num < p.data[j].num }
func (p *dayArray) Swap(i, j int)      { p.data[i], p.data[j] = p.data[j], p.data[i] }

func days() {
	Sunday := day{0, "SUN", "Sunday"}
	Monday := day{1, "MON", "Monday"}
	Tuesday := day{2, "TUE", "Tuesday"}
	Wednesday := day{3, "WED", "Wednesday"}
	Thursday := day{4, "THU", "Thursday"}
	Friday := day{5, "FRI", "Friday"}
	Saturday := day{6, "SAT", "Saturday"}
	// create some sample data
	data := []*day{
		&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday,
	}

	a := dayArray{data}
	sort.Sort(&a)

	if !sort.IsSorted(&a) {
		fmt.Println("Somethine went wrong with IsSorted()")
	}

	// If we get here then all is well
	for _, d := range data {
		fmt.Printf("%s ", d.longName)
	}
	fmt.Println("\n")
}

func main() {
	ints()
	strings()
	days()
}
