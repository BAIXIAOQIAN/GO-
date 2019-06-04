package main

//通过go+匿名函数形式启动goroutine

import (
	"runtime"
	"time"
)

func main() {
	go func() {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
		}
		println(sum)
		time.Sleep(1 * time.Second)
	}()

	//NumGoroutine 可以返回当前程序的 goroutine数目
	println("NumGoroutine = ", runtime.NumGoroutine())

	//main goroutine 故意 "sleep" 5秒，防止其提前退出
	time.Sleep(5 * time.Second)
}
