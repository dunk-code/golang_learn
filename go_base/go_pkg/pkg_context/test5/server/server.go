package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	num := rand.Intn(2)
	if num == 0 {
		time.Sleep(time.Second * 5)
		fmt.Fprint(w, "slow response")
		return
	}
	fmt.Fprint(w, "quick response")
}
