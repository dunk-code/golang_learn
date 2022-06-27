package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func work(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	wg.Done()
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go work(ctx)
	time.Sleep(time.Second * 2)
	cancel()
	wg.Wait()
	fmt.Println("over")
}
