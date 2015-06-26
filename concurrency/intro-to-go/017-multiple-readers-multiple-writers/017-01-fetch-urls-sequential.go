package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Get the page and return the size of a given url and error if any
func getPage(url string) (int, error) {
	resp, err := http.Get(url)

	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return 0, err
	}

	return len(body), nil
}

func main() {

	urls := []string{
		"http://www.google.com",
		"http://www.apple.com",
		"http://www.yahoo.com",
		"http://www.bing.com",
	}

	for _, url := range urls {
		size, err := getPage(url)
		if err != nil {
			os.Exit(1)
		} else {
			fmt.Printf("%s has lenght %d\n", url, size)
		}
	}
}
