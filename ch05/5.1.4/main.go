package main

import (
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup
var urls = []string{
	"http://www.golang.org/",
	"http://www.google.com/",
	"http://www.qq.com",
}

func main() {
	for _, url := range urls {

		//每一个URL，启动一个goroutine，同时给wg加1
		wg.Add(1)

		//Lanch a goroutine to fetch the URL
		go func(url string) {
			//当前goroutine结束后给wg计数减一
			defer wg.Done()

			//发送Http get 请求并打印HTTP返回码
			resp, err := http.Get(url)
			if err == nil {
				println(resp.Status)
			}
		}(url)
	}
	//等待所有请求
	wg.Wait()
	time.Sleep(time.Second * 3)
}
