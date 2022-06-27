package job

import (
	"fmt"
	"time"
)

type EatLunchJob struct {
}

func (j *EatLunchJob) Run() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "结束工作，该吃午饭啦")
}
