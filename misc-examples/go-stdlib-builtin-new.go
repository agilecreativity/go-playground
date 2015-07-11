package main

import "log"

type Actor struct {
	Name string
}

type Movie struct {
	Title  string
	Actors []*Actor
}

func main() {
	ip := new(int)
	log.Printf("ip type: %T, ip: %v, *ip: %v", ip, ip, *ip)

	m := new(Movie)
	log.Printf("m type: %T, m: %v, *m: %v", m, m, *m)
}

/* Output:
2015/07/11 19:42:48 ip type: *int, ip: 0x2081f4410, *ip: 0
2015/07/11 19:42:48 m type: *main.Movie, m: &{ []}, *m: { []}
*/
