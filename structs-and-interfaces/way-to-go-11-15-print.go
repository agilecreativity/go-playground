package main

import "strconv"
import "os"

type Stringer interface {
	String() string
}

type Celcius float64

func (c Celcius) String() string {
	return strconv.FormatFloat(float64(c), 'f', 1, 64) + "C"
}

type Day int

var dayName = []string{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
}

func (day Day) String() string {
	return dayName[day]
}

func print(args ...interface{}) {
	for i, arg := range args {
		if i > 0 {
			os.Stdout.WriteString(" ")
		}

		switch a := arg.(type) {
		case Stringer:
			os.Stdout.WriteString(a.String())
		case int:
			os.Stdout.WriteString(strconv.Itoa(a))
		case string:
			os.Stdout.WriteString(a)
			// TODO: add more type if required!
		default:
			os.Stdout.WriteString("???")
		}
	}
}

func main() {
	print(Day(1), "was", Celcius(18.36))
}

/* Output:
Tuesday was 18.4C
*/
