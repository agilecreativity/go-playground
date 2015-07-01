// Chapter 1: page 25 - 26
package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

var finalString string

func addToFinalStack(letterCh chan string, wg *sync.WaitGroup) {
	letter := <-letterCh
	finalString += letter
	wg.Done()
}

func capitalize(letterCh chan string, currentLetter string, wg *sync.WaitGroup) {
	thisLetter := strings.ToUpper(currentLetter)
	wg.Done()
	letterCh <- thisLetter
}

func main() {
	var wg sync.WaitGroup

	runtime.GOMAXPROCS(2)

	initialString := `
Four score and seven years ago our fathers
brought forth on this continent, a new nation, conceived in
Liberty, and dedicated to the proposition that all men are
created equal.`

	initialBytes := []byte(initialString)

	var letterCh chan string = make(chan string)

	stringLength := len(initialBytes)

	for i := 0; i < stringLength; i++ {
		wg.Add(2)

		go capitalize(letterCh, string(initialBytes[i]), &wg)

		go addToFinalStack(letterCh, &wg)

		wg.Wait()
	}

	fmt.Println(finalString)
}
