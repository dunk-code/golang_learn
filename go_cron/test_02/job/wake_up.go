package job

import (
	"fmt"
	"time"
)

type WakeUpJob struct {
}

func (j *WakeUpJob) Run() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "起床啦...")
}
