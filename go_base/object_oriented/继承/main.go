package main

import "fmt"

type Human struct {
	Name string
	Sex  string
}

func (h *Human) Eat() {
	fmt.Printf("Human:%s eat....\n", h.Name)
}

func (h *Human) Walk() {
	fmt.Printf("Human:%s walk....\n", h.Name)
}

type SuperMan struct {
	Human // 表示SuperMan继承Human
	Level int
}

func main() {
	superman := SuperMan{Human{Name: "zhang3", Sex: "man"}, 30}
	superman.Eat()
}
