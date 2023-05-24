package ejer_interfaces

import (
	"fmt"

	"github.com/go-from-scratch/interfaces"
)

func HumansRespirando(hu interfaces.Human) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s y estoy respirando \n", hu.Sexo())
}
