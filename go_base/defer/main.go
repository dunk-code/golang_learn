package main

import (
	"fmt"
	"io"
	"os"
)

func test(x, y int) {
	var z int
	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()
		panic("divide by zero")
		z = x / y
		return
	}()
	fmt.Printf("x / y = %d\n", z)
}

func main() {
	// test(2, 1)
	do()
}

func do() error {
	f, err := os.Open("book.txt")
	if err != nil {
		return err
	}
	fmt.Println(123)
	if f != nil {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil { // f : another-book.txt
				fmt.Printf("defer close book.txt err %v\n", err)
			}
		}(f)
	}

	// ..code...

	f, err = os.Open("another-book.txt")
	if err != nil {
		return err
	}
	if f != nil {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil { // f : another-book.txt
				fmt.Printf("defer close another-book.txt err %v\n", err)
			}
		}(f)
	}

	return nil
}
