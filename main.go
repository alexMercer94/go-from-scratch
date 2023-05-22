package main

import (
	"fmt"
	"runtime"
)

func main() {
	/* state, text := variables.ComvertToText(1599)
	fmt.Println(state)
	fmt.Println(text) */

	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("No es mi sistema")
	} else {
		fmt.Println("Esto es", os)
	}
}
