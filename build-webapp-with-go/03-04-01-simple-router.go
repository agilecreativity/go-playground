package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Handle the default route
	if r.URL.Path == "/" {
		sayHello(w, r)
		return
	}

	http.NotFound(w, r)
	return
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello router!")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}

/* Note:
try the following in your browser:
http://localhost:9090/    //=> get "Hello router!"
http://localhost:9090/xyz //=> get 404 page not found
*/
