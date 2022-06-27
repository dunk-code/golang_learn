package main

import (
	"fmt"
	"os"
	"time"
)

type PathError struct {
	path       string
	op         string
	createTime string
	message    string
}

func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s, \nop=%s, \ncreateTime=%s, \nmessage=%s", p.path, p.op, p.createTime, p.message)
}

func Open(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		pathError := &PathError{path: filename, op: "read", createTime: fmt.Sprintf("%v", time.Now()), message: err.Error()}
		return pathError
	}
	defer func() {
		file.Close()
	}()
	return nil
}

func main() {
	err := Open("123.txt")
	if err != nil {
		fmt.Println(err)
	}
}
