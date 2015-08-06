package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		cruTime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(cruTime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("04-02-03-duplicate-prevention.gtpl")
		t.Execute(w, token)

	} else {
		// log in request
		r.ParseForm()
		token := r.Form.Get("token")

		if token != "" {
			// check token validity
			fmt.Println("TODO: check if the token is valid: %s\n", token)
		} else {
			// give error if no token
			fmt.Println("TODO: handle error as token is not valid!")
		}

		fmt.Printf("Username length: %v\n", len(r.Form["username"][0]))
		fmt.Printf("Username       : %v\n", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Printf("password       : %v\n", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
	}
}

func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
