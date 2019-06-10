package main

import "runtime"

//缓冲通道和消息队列类似，有削峰和增大吞吐量的功能，示例如下：
func main() {
	c := make(chan struct{})
	c1 := make(chan int, 100)
	go func(i chan struct{}, j chan int) {
		for i := 0; i < 10; i++ {
			c1 <- i
		}
		close(c1)

		//写通道
		c <- struct{}{}
	}(c, c1)

	//NumGoroutine 可以返回当前程序的goroutine数目
	println("NumGoroutine = ", runtime.NumGoroutine())

	//读通道c，通过通道进行同步等待
	<-c

	//此时c1 通道已经关闭，匿名函数启动的goroutine已经退出
	println("NumGoroutine = ", runtime.NumGoroutine())

	//但通道c1还可以继续读取
	for v := range c1 {
		println(v)
	}
}
