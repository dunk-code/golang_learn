package main

import "fmt"

func a() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}
func main() {
	c := a()
	c()
	c()
	c()
	c2 := a()
	c2()
	c2()
	c2()
	a()
}
