package middleware

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a * b
}

func MyMiddleware() {
	fmt.Println("Inicio")
	result := operMid(sumar)(2, 3)
	fmt.Println(result)
	result = operMid(restar)(10, 3)
	fmt.Println(result)
	result = operMid(mult)(2, 4)
	fmt.Println(result)
}

func operMid(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operacion")
		return f(a, b)
	}
}
