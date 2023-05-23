package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num1 int
var num2 int
var message string
var err error

func InputNumbers() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese el primer numero: ")
	if scanner.Scan() {
		num1, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("El dato ingresado es incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese el segundo numero: ")
	if scanner.Scan() {
		num2, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("El dato ingresado es incorrecto" + err.Error())
		}
	}

	fmt.Println("Ingrese mensaje: ")
	if scanner.Scan() {
		message = scanner.Text()
	}

	fmt.Println(message, num1*num2)
}
