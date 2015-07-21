package main

import (
	"encoding/asn1"
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
	}
}

func main() {
	mdata, err := asn1.Marshal(13)

	checkError(err)

	fmt.Printf("After marshal: %#v\n", mdata)

	var n int

	_, err = asn1.Unmarshal(mdata, &n)

	checkError(err)

	fmt.Printf("After unmarshal: %#v\n", n)
}

/*
After marshal: []byte{0x2, 0x1, 0xd}
After unmarshal: 13
*/
