package main

import (
	"fmt"
	"runtime"
)

func main() {
	/* state, text := variables.ComvertToText(1599)
	fmt.Println(state)
	fmt.Println(text) */

	if os := runtime.GOOS; os == "linux" || os == "OS X" {
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
	}
}
