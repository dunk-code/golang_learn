package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for true {
		fmt.Println("worker")
		time.Sleep(time.Second * 1)
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			return
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	go worker(ctx)
	time.Sleep(time.Second * 2)
	cancel()
	fmt.Println("over")
}
