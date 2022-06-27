package main

import "fmt"

func fuck(c chan int) {
	data := <-c
	fmt.Println(data)
}

func main() {
	c := make(chan int)
	go fuck(c)
	c <- 13
	x, y := 1, 3
	swap(&x, &y)
	fmt.Printf("x = %d, y = %d\n", x, y)
}

func swap(x, y *int) {
	fmt.Println(x, y)
	*x, *y = *y, *x
	fmt.Println(x, y)
}
