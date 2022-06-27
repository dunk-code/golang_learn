package main

import (
	"context"
	"fmt"
)

func work(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for true {
			select {
			case <-ctx.Done():
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	dst := work(ctx)
	for n := range dst {
		fmt.Println(n)
		if n == 5 {
			return
		}
	}
}
