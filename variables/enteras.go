package variables

import "fmt"

func ShowInts() {
	var intcomun int
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("intcomun =", intcomun)
	fmt.Println("intde32 =", intde32)
	fmt.Println("intde64 =", intde64)
}
