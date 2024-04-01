package main

import "fmt"

// 定义函数fib, 实现斐波那契
func fib(n int, c chan int) {
	x, y := 1, 1

	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	// 显式关闭管道
	// 当管道被关闭后将不能再向管道写入数据，但还可以从管道读取数据。
	close(c)
}
func main() {
	fmt.Println(111)
	c := make(chan int, 10)

	go fib(cap(c), c)

	// for i := range c 能够不断地读取channel里面的数据，main
	// 直到该channel被显式关闭
	for i := range c {
		fmt.Println(i)
	}
}
