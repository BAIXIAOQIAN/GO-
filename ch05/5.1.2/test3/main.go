package main

/*func GOMAXPROCS
  func GOMAXPROCS(n int)用来设置或查询可以并发执行的goroutine数目，n大于1表示设置GOMAXPROCS值，否则表示查询当前的GOMAXPROCS值
*/

import "runtime"

func main() {

	//获取当前的GOMAXPROCS值
	println("GOMAXPROCS = ", runtime.GOMAXPROCS(0))

	//设置当前的GOMAXPROCS的值为2
	runtime.GOMAXPROCS(2)

	//获取当前的GOMAXPROCS值
	println("GOMAXPROCS = ", runtime.GOMAXPROCS(0))
}
