package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}

	url := os.Args[1]

	response, err := http.Head(url)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	fmt.Println(response.Status)
	for k, v := range response.Header {
		fmt.Println(k+":", v)
	}

	os.Exit(0)
}

/* output:
$./08-03-01-head http://www.google.com:80

200 OK
X-Frame-Options: [SAMEORIGIN]
Set-Cookie: [PREF=ID=1111111111111111:FF=0:TM=1437823505:LM=1437823505:V=1:S=M8xoREmzL7NLdKp2; expires=Mon, 24-Jul-2017 11:25:05 GMT; path=/; domain=.google.com.au NID=69=UIr0rKaOXVaIugiFPMkVfhki9keMTcPgNTA0W58hM4zAdbGzu4bWZYKXmMDKskLoelsrt17lgNU3Udt8FLDmDFjrLYGkDIAZK-wbzQIN1XmGSksIQW36sZI3JYwnRfH_; expires=Sun, 24-Jan-2016 11:25:05 GMT; path=/; domain=.google.com.au; HttpOnly]
Expires: [-1]
P3p: [CP="This is not a P3P policy! See http://www.google.com/support/accounts/bin/answer.py?hl=en&answer=151657 for more info."]
Alternate-Protocol: [80:quic,p=0]
Accept-Ranges: [none]
Date: [Sat, 25 Jul 2015 11:25:05 GMT]
Cache-Control: [private, max-age=0]
Content-Type: [text/html; charset=ISO-8859-1]
Vary: [Accept-Encoding]
Server: [gws]
X-Xss-Protection: [1; mode=block]
*/
