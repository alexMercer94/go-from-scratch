package funciones

import "fmt"

/*
Funciones anonimas
*/
func Calculos() {
	var num3 int = 32
	var num33 int = 434

	calc := func(number1 int, number2 int) int {
		return number1 + number2 + num3 + num33
	}

	fmt.Println(calc(20, 24))
	calc = func(number1 int, number2 int) int {
		return number1 * number2 * num3
	}

	fmt.Println(calc(20, 24))
}
