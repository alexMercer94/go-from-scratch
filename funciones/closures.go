package funciones

import "fmt"

func table(value int) func() int {
	number := value
	sec := 0

	return func() int {
		sec++
		return number * sec
	}
}

func CallClosure() {
	tabledel := 2
	MiTabla := table(tabledel)
	for i := 1; i <= 10; i++ {
		fmt.Println(MiTabla())
	}
}
