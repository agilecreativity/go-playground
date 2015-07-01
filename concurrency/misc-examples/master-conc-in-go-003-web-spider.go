// Chapter 1: page 35
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var applicationStatus bool
var urls []string
var urlsProcessed int
var foundUrls []string
var fullText string
var totalUrlCount int
var wg sync.WaitGroup

var v1 int

func readUrls(statCh chan int, textCh chan string) {
	time.Sleep(1 * time.Millisecond)
	fmt.Println("Grabbing", len(urls), "urls")

	// TODO: remove the global var
	for i := 0; i < totalUrlCount; i++ {
		fmt.Println("Url", i, urls[i])

		resp, _ := http.Get(urls[i])

		text, err := ioutil.ReadAll(resp.Body)
		xxx := string(text)
		fmt.Printf("FYI: %s\n", xxx)
		textCh <- xxx // string(text)

		if err != nil {
			fmt.Println("No HTML body")
		}

		statCh <- 0
	}
}

func addToScrapedText(textCh chan string, procCh chan bool) {
	for {
		select {
		case pc := <-procCh:
			if pc {
				// hang on
			} else {
				close(textCh)
				close(procCh)
			}
		case text := <-textCh:
			// TODO: global varaiable!
			fullText += text
		}
	}
}

func evaluateStatus(statCh chan int, textCh chan string, procCh chan bool) {
	for {
		select {
		case status := <-statCh:
			fmt.Print(urlsProcessed, totalUrlCount)
			urlsProcessed++

			// TODO: clean this up
			if status == 0 {
				fmt.Println("Got url...")
			}
			if status == 1 {
				close(statCh)
			}
			if urlsProcessed == totalUrlCount {
				fmt.Println("Read all top-level Urls")
				procCh <- false
				applicationStatus = false
			}
		}
	}
}

func main() {
	applicationStatus = true
	statCh := make(chan int)
	textCh := make(chan string)
	procCh := make(chan bool)
	totalUrlCount = 0

	urls = []string{
		"http://www.google.com",
		"http://www.yahoo.com",
	}

	fmt.Println("Starting spider..")

	urlsProcessed = 0
	totalUrlCount = len(urls)

	go evaluateStatus(statCh, textCh, procCh)
	go readUrls(statCh, textCh)
	go addToScrapedText(textCh, procCh)

	for {
		if !applicationStatus {
			fmt.Println(fullText)
			fmt.Println("Done")
			break
		}
		select {
		case status := <-statCh:
			fmt.Println("Message on status channel", status)
		}
	}
}
