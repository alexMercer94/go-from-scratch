package defpan

import (
	"fmt"
	"log"
)

func ShowDefer() {
	fmt.Println("Estes es un primer mensaje")
	defer fmt.Println("Este es el mensaje final")

	fmt.Println("Estes es el segundo mensaje")
}

func EjempPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que genero Panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	}
}
