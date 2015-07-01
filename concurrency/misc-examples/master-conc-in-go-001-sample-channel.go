// chapter 1: page 23
package main

import (
	"fmt"
	"runtime"
	"strings"
)

func deliverToFinal(letter string, finalIpsum *string) {
	*finalIpsum += letter
}

func capitalize(current *int, length int, letters []byte, finalIpsum *string) {
	for *current < length {
		thisLetter := strings.ToUpper(string(letters[*current]))
		deliverToFinal(thisLetter, finalIpsum)
		*current++
	}
}

func main() {
	var loremIpsum string
	var finalIpsum string

	runtime.GOMAXPROCS(2)
	index := new(int)
	*index = 0
	loremIpsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum venenatis magna eget libero tincidunt, ac condimentum enim auctor. Integer mauris arcu, dignissim sit amet convallis vitae, ornare vel odio. Phasellus in lectus risus. Ut sodales vehicula ligula eu ultricies. Fusce vulputate fringilla eros at congue. Nulla tempor neque enim, non malesuada arcu laoreet quis. Aliquam eget magna metus. Vivamus lacinia venenatis dolor, blandit faucibus mi iaculis quis. Vestibulum sit amet feugiat ante, eu porta justo."
	letters := []byte(loremIpsum)
	length := len(loremIpsum)
	go capitalize(index, length, letters, &finalIpsum)

	go func() {
		go capitalize(index, length, letters, &finalIpsum)
	}()

	fmt.Println(length, " characters.")
	fmt.Println(loremIpsum)
	fmt.Println(*index)
	fmt.Println(finalIpsum)
}
