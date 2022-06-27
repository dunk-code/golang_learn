package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("*/7 * * * * *", func() {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	})
	c.Start()
	time.Sleep(300 * time.Second)
}
