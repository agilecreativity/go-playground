package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Task struct {
	value      int
	executedBy string
}

var total int
var task Task
var lock sync.Mutex

func calculate() {
	for total < 10 {
		genRandom("from calculate()")
	}
}

func perform() {
	for total < 10 {
		genRandom("from perform()")
	}
}

func genRandom(srcName string) {
	lock.Lock()
	task.value = rand.Intn(100)
	task.executedBy = srcName
	display()
	total++
	lock.Unlock()
	time.Sleep(500 * time.Millisecond)
}

func display() {
	fmt.Println("----------")
	fmt.Println(task.value)
	fmt.Println(task.executedBy)
	fmt.Println("----------")
}

func main() {
	fmt.Println("Synchronizing goroutines demo")

	go calculate()
	go perform()

	var input string
	fmt.Scanln(&input)
	fmt.Println("done!")
}
