// Use empty struct and assign the value
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type webPage struct {
	url  string
	body []byte
	err  error
}

func (w *webPage) get() {
	resp, err := http.Get(w.url)

	if err != nil {
		w.err = err
		return
	}

	defer resp.Body.Close()
	w.body, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		w.err = err
	}
}

func (w *webPage) isOK() bool {
	return w.err == nil
}

func main() {
	url := "http://www.oreilly.com"
	// empty construct
	w := &webPage{}
	w.url = url
	w.get()

	if !w.isOK() {
		fmt.Printf("Error fetching %s : %s", url, w.err)
		return
	}

	fmt.Printf("URL: %s, Length: %d", w.url, len(w.body))
}
