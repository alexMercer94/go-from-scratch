package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num1 int
var err error

func CreateTable() string {
	scanner := bufio.NewScanner(os.Stdin)
	var table string

	for {
		fmt.Println("Ingrese el numero: ")
		if scanner.Scan() {
			num1, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	fmt.Printf("\n")
	for i := 1; i <= 10; i++ {
		table += fmt.Sprintf("%d x %d = %d \n", num1, i, num1*i)
	}

	return table
}
