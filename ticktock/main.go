package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10000; i++ {
		if i%2 == 0 {
			fmt.Printf("Tick %d\n", i)
		} else {
			fmt.Printf("Tock %d\n", i)
		}
		time.Sleep(1000 * time.Millisecond)
	}
}
