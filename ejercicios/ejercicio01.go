package ejercicios

import (
	"fmt"
	"strconv"
)

func ReturnValues(parameter string) (int, string) {
	value, err := strconv.Atoi(parameter)
	var message string

	if err != nil {
		fmt.Println("Error during conversion")
		return 0, "Error"
	}

	if value > 100 {
		message = "Es mayor a 100"
	} else {
		message = "Es menor a 100"
	}

	return value, message
}
