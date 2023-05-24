package main

import (
	e "github.com/go-from-scratch/ejer_interfaces"
	"github.com/go-from-scratch/models"
)

func main() {
	/* state, text := variables.ComvertToText(1599)
	fmt.Println(state)
	fmt.Println(text) */

	// Condicionales
	/* if os := runtime.GOOS; os == "linux" || os == "OS X" {
		fmt.Println("No es mi sistema, es:", os)
	} else {
		fmt.Println("Esto es Mac", os)
	}

	switch os := runtime.GOOS; os {
	case "Linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es MacOS")
	default:
		fmt.Printf("%s \n", os)
	} */

	// Ejercicio 01
	/* number, message := ejercicios.ReturnValues("55")
	fmt.Println("Number: ", number)
	fmt.Println("Message: ", message) */

	// teclado.InputNumbers()

	// iteraciones.Iterate()

	// Ejercicio 02
	// fmt.Println(ejercicios.CreateTable())

	// files.SaveTable()
	// files.SumTableInFile()
	// files.ReadFile()

	// funciones.Exponencia(2)

	// Arrays
	// arrays_slices.Capacidad()

	// Mapas
	// mapas.ShowMapas()

	// Objectos POO
	// users.AddUser()

	// Interfaces
	Pedro := new(models.Hombre)
	e.HumansRespirando(Pedro)
	Maria := new(models.Mujer)
	e.HumansRespirando(Maria)
}
