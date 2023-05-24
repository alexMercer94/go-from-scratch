package mapas

import "fmt"

func ShowMapas() {
	paises := make(map[string]string)
	// fmt.Println(paises)

	paises["Mexico"] = "DF"
	paises["Argentina"] = "Buenos Aires"

	/* fmt.Println(paises)
	fmt.Println(paises["Argentina"]) */

	camp := map[string]int{
		"Barcelona":   39,
		"Real Madrid": 38,
		"America":     37,
		"Boca":        30,
	}

	fmt.Println(camp)

	for equipo, puntaje := range camp {
		fmt.Printf("Equipp %s, tiene un puntaje de %d \n", equipo, puntaje)
	}

	delete(camp, "Barcelona")
	fmt.Println(camp)

	puntaje, exist := camp["America"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t", puntaje, exist)
}
