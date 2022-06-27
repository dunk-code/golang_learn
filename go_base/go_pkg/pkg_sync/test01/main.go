package main

import (
	"fmt"
	"time"
)

func main() {
	var a = 0
	for i := 0; i < 100; i++ {
		go func(idx int) {
			a += 1
			fmt.Printf("goroutine %d, a=%d\n", idx, a)
		}(i)
	}
	time.Sleep(time.Second * 1)
}
