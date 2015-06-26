package main

import "fmt"
import "math/rand"
import "time"

func main() {
	boring("boring!")
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
