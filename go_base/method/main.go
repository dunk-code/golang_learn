package main

import "fmt"

type T struct {
	int
}

func (t T) testT() {
	fmt.Println("TestT")
}

func (t *T) testP() {
	fmt.Println("TestP")
}

type S struct {
	T
}

func main() {
	t1 := T{1}
	t1.testT()
	t2 := &t1
	t2.testT()
	t2.testP()
	s1 := S{t1}
	s1.testT()
	s2 := &s1
	s2.testT()

}
