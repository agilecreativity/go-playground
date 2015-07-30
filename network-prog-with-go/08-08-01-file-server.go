package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// deliver files from the directory /var/www
	// fileServer := http.FileServer(http.Dir("/var/www"))

	// Note: adjust to your local settings
	fileServer := http.FileServer(http.Dir("/tmp"))

	// register the handler and deliver requests to it
	err := http.ListenAndServe(":8000", fileServer)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
