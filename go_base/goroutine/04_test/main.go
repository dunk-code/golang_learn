package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		defer fmt.Println("defer A")
		func() {
			defer fmt.Println("defer B")
			runtime.Goexit()
			defer fmt.Println("defer C")
			fmt.Println("A")
		}()
		fmt.Println("B")
	}()
	for {
	}
}
