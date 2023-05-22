package main

import (
	"fmt"

	"github.com/go-from-scratch/variables"
)

func main() {
	state, text := variables.ComvertToText(1599)
	fmt.Println(state)
	fmt.Println(text)
}
