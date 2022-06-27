package main

import (
	"github.com/robfig/cron/v3"
	"golang_learn/go_cron/test_02/job"
)

func main() {
	c := cron.New(cron.WithSeconds())
	wakeUpJob := &job.WakeUpJob{}
	c.AddJob("0 55 8 * * 1-5", wakeUpJob) // 周一到周五每天8:55触发一次
	lunchJob := &job.EatLunchJob{}
	c.AddJob("0 0 12 * * *", lunchJob) // 每天12:00触发一次
	c.Start()
}
