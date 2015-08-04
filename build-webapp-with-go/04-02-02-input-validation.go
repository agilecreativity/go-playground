package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
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
		// Use our simple template file
		t, _ := template.ParseFiles("04-02-01-input-validation.gtpl")
		t.Execute(w, nil)
	} else {
		// Handle the "POST" method
		r.ParseForm()
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])

		// Another way to get the form parameter
		age := r.Form.Get("age")
		fmt.Println("age :", age)

		// Let's validate the input
		err := validateAge(age)
		if err != nil {
			// Note: just log the validation for now!
			fmt.Printf("FYI: fail validation :%s", err.Error())
		}
	}
}

// Check and validate the age input
func validateAge(age string) (err error) {
	ageValue, err := strconv.Atoi(age)
	if ageValue > 100 {
		err = errors.New("You are too old!")
	}

	// Or using regular expression
	re := regexp.MustCompile("^[0-9]+$")

	if !re.MatchString(age) {
		err = errors.New("Not a number :" + age)
	}

	return
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
// This will be handled by the POST request
http://localhost:9090/

// These will be handled by the 'GET' request
http://localhost:9090/login?user_name=john&password=topsecret&age=120
http://localhost:9090/login?user_name=john&password=topsecret&age=not-number
*/
