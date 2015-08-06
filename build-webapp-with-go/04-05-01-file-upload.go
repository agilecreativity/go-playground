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
		r.ParseMultipartForm(32 << 20)
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
	http.HandleFunc("/", upload)
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

/* Try:
http://localhost:9090/long_url
http://localhost:9090/login?user_name=john&password=topsecret
*/
