package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Name string
var State bool
var Sueldo float32
var Date time.Time

func RestVars() {
	Name := "Alex"
	State := true
	Sueldo := 1740.66
	Date := time.Now()

	fmt.Println(Name)
	fmt.Println(State)
	fmt.Println(Sueldo)
	fmt.Println(Date)
}

func ComvertToText(number int) (bool, string) {
	text := strconv.Itoa(number)
	return true, text

}
