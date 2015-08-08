// TODO: complete this example!
package main

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type RestInterface interface {
	Get(values url.Values) (int, interface{})
	Post(values url.Values) (int, interface{})
	Put(values url.Values) (int, interface{})
	Delete(values url.Values) (int, interface{})
}

func RestController(i RestInterface) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		values := r.Form
		method := r.Method
		switch method {
		case "GET":
			code, data := i.Get(values)
		case "POST":
			code, data := i.Post(values)
		case "PUT":
			code, data := i.Put(values)
			// ignore "DELETE" for now
		default:
			code, data := 405, nil
		}

		j, err := json.Marshal(data)
		if err != nil {
			code = 500 // internal server error
		}

		if code >= 400 {
			http.Error(rw, http.StatusText(code), code)
			return
		}

		rw.WriteHeader(code)
		rw.Write(j)
	}
}

func main() {
	// TODO:
}
