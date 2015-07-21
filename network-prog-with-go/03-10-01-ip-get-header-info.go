package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte

	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage  : %s host:port \n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example: %s www.google.com:80 \n", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	conn, err := net.Dial("tcp", service)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := readFully(conn)
	checkError(err)

	fmt.Println(string(result))
	os.Exit(0)
}

/* Output:
$./03-10-01-ip-get-header-info www.google.com:80

HTTP/1.0 302 Found
Cache-Control: private
Content-Type: text/html; charset=UTF-8
Location: http://www.google.com.au/?gfe_rd=cr&ei=sh2uVf7OOMXu8wep9rfQCQ
Content-Length: 262
Date: Tue, 21 Jul 2015 10:23:46 GMT
Server: GFE/2.0
Alternate-Protocol: 80:quic,p=0


*/
