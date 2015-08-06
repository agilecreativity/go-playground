package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/textproto"
	"os"
	"strconv"
	"text/template"
	"time"
)

type FileHeaer struct {
	Filename string
	Header   textproto.MIMEHeader
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		cruTime := time.Now().Unix()
		h := md5.New()

		io.WriteString(h, strconv.FormatInt(cruTime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("04-05-01-file-upload.gtpl")
		t.Execute(w, token)
	} else {
		// The parameter is maxMemory int64
		// See: https://golang.org/pkg/net/http/#Request.ParseMultipartForm
		const maxMemory int64 = 32 << 20
		fmt.Printf("FYI: maxMemory used :%d\n", maxMemory)
		r.ParseMultipartForm(maxMemory)
		file, handler, err := r.FormFile("uploadfile")

		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		// Note: need to create the sample test folder
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = ":9090"
	}

	http.HandleFunc("/", upload)
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
