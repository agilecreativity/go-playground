package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// Perform post of a given file without the UI
func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// create the file that will be uploaded!
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)

	if err != nil {
		fmt.Println("Error writing to buffer")
		return err
	}

	// open file handle
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", filename)
		return err
	}

	// iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	// save it for later
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	// Post the file
	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	fmt.Printf("Response Status: %s\n", resp.Status)
	fmt.Printf("Response Body  : %s\n", resp_body)
	return nil
}

func main() {
	// Note: make sure that you start the server in 04-05-01-upload.go
	targetUrl := "http://localhost:9090/upload"
	fileName := "./04-05-01-file-upload.gtpl"
	postFile(fileName, targetUrl)
}
