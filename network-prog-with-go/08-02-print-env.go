package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// TODO: adjust to your own environment
	fileServer := http.FileServer(http.Dir("/Users/bchoomnuan/codes/spikes/"))
	http.Handle("/", fileServer)
	http.HandleFunc("/cgi-bin/printenv", printEnv)
	err := http.ListenAndServe(":8000", nil)
	checkError(err)
}

func printEnv(writer http.ResponseWriter, req *http.Request) {
	env := os.Environ()

	writer.Write([]byte("<h1>Environment</h1>\n<pre>"))

	for _, v := range env {
		writer.Write([]byte(v + "\n"))
	}

	writer.Write([]byte("</pre>"))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
