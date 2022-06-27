package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		fmt.Println(1)
		fmt.Println(2)
		fmt.Println(3)
		wg.Done()
	}()

	go func() {
		fmt.Println(12)
		fmt.Println(13)
		fmt.Println(14)
		wg.Done()
	}()

	go func() {
		fmt.Println(22)
		fmt.Println(33)
		time.Sleep(time.Second)
		fmt.Println(44)
		wg.Done()
	}()

	wg.Wait()
}
