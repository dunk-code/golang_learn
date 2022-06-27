package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func main() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s + strconv.Itoa(i))
		}
	}("goroutine")
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println("main")
	}
}
