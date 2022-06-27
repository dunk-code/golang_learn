package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for true {
			i, ok := <-ch1
			if ok {
				ch2 <- i
			} else {
				break
			}
		}
		close(ch2)
	}()
	for i := range ch2 {
		fmt.Println(i)
	}
}
