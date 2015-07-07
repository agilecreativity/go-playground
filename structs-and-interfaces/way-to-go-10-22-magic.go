package main

import "fmt"

type Base struct{}

func (Base) Magic() {
	fmt.Print("base magic ")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

func main() {
	v := new(Voodoo)
	v.Magic()     //=> voodo magic
	v.MoreMagic() //=> base magic base magic
}

/* Output:
voodoo magic
base magic base magic
*/
