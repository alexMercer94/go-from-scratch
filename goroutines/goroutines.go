package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MyNameSlowly(name string, canal1 chan bool) {
	letters := strings.Split(name, "")

	for _, letter := range letters {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letter)
	}

	canal1 <- true
}
