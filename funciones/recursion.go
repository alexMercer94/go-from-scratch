package funciones

import "fmt"

func Exponencia(value int) {
	if value > 10000 {
		return
	}

	fmt.Println(value)
	Exponencia(value * 2)
}
