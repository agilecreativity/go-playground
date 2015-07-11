package main

import "log"

func main() {
	m := make(map[string]int)
	log.Println(m)

	m["one"] = 1
	log.Println(m)

	m["two"] = 2
	log.Println(m)

	delete(m, "one")
	log.Println(m)
	m = nil
	delete(m, "two")

	log.Println(m)
}

/* Output:
2015/07/11 19:46:01 map[]
2015/07/11 19:46:01 map[one:1]
2015/07/11 19:46:01 map[one:1 two:2]
2015/07/11 19:46:01 map[two:2]
2015/07/11 19:46:01 map[]
*/
