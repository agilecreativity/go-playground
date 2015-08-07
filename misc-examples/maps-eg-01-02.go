package main

import "fmt"

type Saiyan struct {
	Name    string
	Friends map[string]*Saiyan
}

func main() {
	goku := &Saiyan{
		Name:    "Goku",
		Friends: make(map[string]*Saiyan),
	}

	goku.Friends["krillin"] = &Saiyan{Name: "xxx", Friends: nil}

	lookup := map[string]int{
		"goku":  9001,
		"gohan": 2044,
	}

	for key, value := range lookup {
		fmt.Printf("%v => %v\n", key, value)
	}
}
