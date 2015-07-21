package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
func main() {
	// port
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue // ignore error
		}
		daytime := time.Now().String()

		// don't care about return value
		conn.Write([]byte(daytime))

		// we're finished with this client
		conn.Close()
	}
}

/* Note: to test just try the following command from terminal
$telnet localhost 1200
# which will produced
telnet localhost 1200
Trying ::1...
Connected to localhost.
Escape character is '^]'.
2015-07-21 19:13:23.489997977 +1000 AESTConnection closed by foreign host.
*/
