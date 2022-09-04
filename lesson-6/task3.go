package main

import (
	"fmt"
	"time"
)

func main() {
	var x int
	for i := 1; i <= 1000; i++ {
		go func() {
			x += 1 // на этой строке обнаружено состояние гонки
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("x =", x)
}
