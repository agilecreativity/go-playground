package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}

	url := os.Args[1]

	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	b, _ := httputil.DumpResponse(response, false)
	fmt.Print(string(b))

	contentTypes := response.Header["Content-Type"]

	if !acceptableCharset(contentTypes) {
		fmt.Println("Cannot handle", contentTypes)
		os.Exit(4)
	}

	var buf [512]byte
	reader := response.Body

	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			os.Exit(0)
		}
		fmt.Print(string(buf[0:n]))
	}
	os.Exit(0)
}

func acceptableCharset(contentTypes []string) bool {
	for _, cType := range contentTypes {
		if strings.Index(cType, "UTF-8") != -1 {
			return true
		}
	}
	return false
}

/* output:
./08-03-02-get http://www.google.com:80

HTTP/1.1 200 OK
Transfer-Encoding: chunked
Accept-Ranges: none
Alternate-Protocol: 80:quic,p=0
Cache-Control: private, max-age=0
Content-Type: text/html; charset=ISO-8859-1
Date: Sat, 25 Jul 2015 11:37:23 GMT
Expires: -1
P3p: CP="This is not a P3P policy! See http://www.google.com/support/accounts/bin/answer.py?hl=en&answer=151657 for more info."
Server: gws
Set-Cookie: PREF=ID=1111111111111111:FF=0:TM=1437824243:LM=1437824243:V=1:S=atYnHd-Gq7qNmm0Y; expires=Mon, 24-Jul-2017 11:37:23 GMT; path=/; domain=.google.com.au
Set-Cookie: NID=69=MAu_pj6zbD3ijy9MBDBeHrefbbD9YiAOtuhNyhfkq1nvy2OB9Z7fiW-86cTXP4Mbx1KRJ_78tAQrd7YZHcTSuaSQ4ghjWp_BBWpX-yWH0N88VP9-EKr6p44sBXK6BZZO; expires=Sun, 24-Jan-2016 11:37:23 GMT; path=/; domain=.google.com.au; HttpOnly
Vary: Accept-Encoding
X-Frame-Options: SAMEORIGIN
X-Xss-Protection: 1; mode=block

Cannot handle [text/html; charset=ISO-8859-1]
*/
