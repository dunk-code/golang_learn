package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type RspData struct {
	rsp *http.Response
	err error
}

func doCall(ctx context.Context) {
	transport := http.Transport{DisableKeepAlives: true}
	client := http.Client{Transport: &transport}
	req, err := http.NewRequest("GET", "http://127.0.0.1:8000/", nil)
	if err != nil {
		panic(err)
	}
	req = req.WithContext(ctx)
	dataChan := make(chan *RspData, 1)
	var wg sync.WaitGroup
	defer wg.Wait()
	doRequest := func() {
		wg.Add(1)
		rsp, err := client.Do(req)
		fmt.Printf("client.do resp:%v, err:%v\n", rsp, err)
		data := &RspData{rsp: rsp, err: err}
		dataChan <- data
		wg.Done()
	}
	go doRequest()
	select {
	case <-ctx.Done():
		fmt.Println("call api timeout")
	case result := <-dataChan:
		fmt.Println("call server api success")
		if result.err != nil {
			fmt.Printf("call server api failed, err:%v\n", result.err)
			return
		}
		defer result.rsp.Body.Close()
		data, _ := ioutil.ReadAll(result.rsp.Body)
		fmt.Printf("resp:%v\n", string(data))
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()
	doCall(ctx)
}
