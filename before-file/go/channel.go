// package main

// import "fmt"

// // 定义函数fib, 实现斐波那契
// func fib(n int, c chan int) {
// 	x, y := 1, 1

// 	for i := 0; i < n; i++ {
// 		c <- x
// 		x, y = y, x+y
// 	}

// 	// 显式关闭管道
// 	// 当管道被关闭后将不能再向管道写入数据，但还可以从管道读取数据。
// 	close(c)
// }
// func main() {
// 	fmt.Println(111)
// 	c := make(chan int, 10)

// 	go fib(cap(c), c)

// 	// for i := range c 能够不断地读取channel里面的数据，main
// 	// 直到该channel被显式关闭
// 	for i := range c {
// 		fmt.Println(i)
// 	}
// }

// select 关键字
// 作用：通过select 可以监听 channel 上的数据流动

// select 默认是阻塞的，只有当监听的channel 中有发送或接收可以进行时才会发生，当多个channel 都准备好的时候，select 是随机的选择一个执行的。
package main

import "fmt"

func fib(c, q chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-q:
			fmt.Println("q channel output")

			// default:
			// 当所有channel都没有准备好，发生阻塞时执行
			return
		}
	}
}

func main() {
	c := make(chan int)

	q := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}

		q <- 0
	}()

	fib(c, q)
}
