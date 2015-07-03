// Page: 142:
// Performing actions in the background
package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for _, v := range os.Args {
		wg.Add(1)
		go func(str string) {
			fmt.Printf("%s\n", strings.ToUpper(str))
			wg.Done()
		}(v)
	}
	wg.Wait()
}
