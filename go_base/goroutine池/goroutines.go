package goroutine池

import (
	"fmt"
	"math/rand"
)

type Job struct {
	Id      int
	RandNum int
}

type Result struct {
	job *Job
	sum int
}

func work() {
	jobChan := make(chan *Job)
	resultChan := make(chan *Result)
	createPool(64, jobChan, resultChan)
	go func(resultChan chan *Result) {
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id,
				result.job.RandNum, result.sum)
		}
	}(resultChan)
	var id int
	for true {
		id++
		jobChan <- &Job{
			Id:      id,
			RandNum: rand.Int(),
		}
	}
}

func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	// 	开启num个协程运行
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			for job := range jobChan {
				r_num := job.RandNum
				var sum int
				for r_num != 0 {
					sum += r_num % 10
					r_num /= 10
				}
				resultChan <- &Result{
					job: job,
					sum: sum,
				}
			}
		}(jobChan, resultChan)
	}
}
