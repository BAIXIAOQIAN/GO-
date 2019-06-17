package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

//通过一个随机数生成器来演示通知退出机制，下游的消费者不需要随机数时，显示的通知生产者停止生产

//GenerrateIntA是一个随机数生成器
func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
				//增加一路监听，就是对退出通知信号done的监听
			case <-done:
				break Lable
			}
		}
		//收到通知后关闭通道ch
		close(ch)
	}()
	return ch
}

func main() {
	done := make(chan struct{})
	ch := GenerateIntA(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	//发出通知，告诉生产者停止生产
	close(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	//此时生产者已经退出
	println("NumGoroutine = ", runtime.NumGoroutine())
}
