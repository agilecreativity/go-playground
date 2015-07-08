package main

import "os"

var user = os.Getenv("USER")

// Note: this could be checked in an init() function of a package that uses this code
func check() {
	if user == "" {
		panic("Unknown user: no value for $USER")
	}
}

func main() {

}
