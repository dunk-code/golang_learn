package main

import "fmt"

type Writer interface {
	write()
}

type Reader interface {
	read()
}

type Book struct {
}

func (b *Book) read() {
	fmt.Println("book read...")
}

func (b *Book) write() {
	fmt.Println("book write...")
}

func main() {
	// b:pair<type:Book, value:Book{}地址>
	b := &Book{}
	var r Reader
	// r:pair<type:Book, value:Book{}地址>
	r = b
	r.read()
	// w:pair<type:Book, value:Book{}地址>
	w := r.(Writer)
	w.write()
}
