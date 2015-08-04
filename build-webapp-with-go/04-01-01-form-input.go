package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r.Form) // print information on server side
	fmt.Println("path  : ", r.URL.Path)
	fmt.Println("scheme: ", r.URL.Scheme)
	fmt.Println(r.Form["long_url"])
	for k, v := range r.Form {
		fmt.Printf("key: %s\n", k)
		fmt.Printf("val: %s\n", strings.Join(v, ""))
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		// use our simple template file
		t, _ := template.ParseFiles("04_01_01_loging.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

/* Try:
http://localhost:9090/long_url
http://localhost:9090/login?user_name=john&password=topsecret
*/
