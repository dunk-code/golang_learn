package main

import (
	"strconv"
	"time"
)

func main() {
	print(strconv.FormatInt(time.Now().Unix(), 10))
}
