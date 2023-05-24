package arrays_slices

import (
	"fmt"
)

var tableS []int = []int{2, 5, 4}
var array [10]int = [10]int{6, 98, 21, 398, 398, 99, 2}

func ShowSlices() {
	fmt.Println(tableS)

	// Slice creado desde un vector desde la posicion 3
	porcion := array[3:]

	// Slice creado desde un vector, desde la posicion 0 hasta la 5
	porcion2 := array[:5]

	// Slice creado desde un vector, desde la posicion 6 hasta la 7
	porcion3 := array[6:7]

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)

}

func Capacidad() {
	// elements := make([]int, 5, 20)

	/* fmt.Printf("Largo %d, Capacidad %d", len(elements), cap(elements))
	 */
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("Largo %d, Capacidad %d", len(nums), cap(nums))
}
