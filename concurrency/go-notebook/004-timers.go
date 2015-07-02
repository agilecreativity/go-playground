// See: page 117
package main

import "time"
import "fmt"

type Clock struct {
	Period  int64
	Count   chan int64
	Control chan bool
	active  bool
}

func (c *Clock) Start() {
	if !c.active {
		go func() {
			c.active = true
			for i := int64(0); ; i++ {
				select {
				case c.active = <-c.Control:
				default:
					if c.active {
						c.Count <- i
					}
					time.Sleep(time.Duration(c.Period))
				}
			}
		}()
	}
}

func main() {
	c := Clock{1000, make(chan int64), make(chan bool), false}
	c.Start()

	for i := 0; i < 3; i++ {
		fmt.Println("Pulse value", <-c.Count, "from clock")
	}

	fmt.Println("disabling clock")

	c.Control <- false
	time.Sleep(time.Duration(1 * time.Microsecond))
	fmt.Println("Restarting clock")
	c.Control <- true

	fmt.Println("Pulse value", <-c.Count, "from clock")
}
