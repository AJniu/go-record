package main

import (
	"fmt"
	"os"
)

var user = os.Getenv("USER")

// panic(): 中断当前代码的执行，并返回错误信息。

// recover(): 捕获当前代码中 panic() 发生的错误，返回错误，并恢复代码执行。
// recover() 必须与 defer 配合使用，否则将只会返回 nil。

func hasPanic() {
	if user == "" {
		// 使用panic中断代码执行，并抛出错误。
		panic("no value for $USER")

		// 以下代码无法执行
		// fmt.Println("no run")
	}
}

func throwsPanic(f func()) (b bool) {
	// 异常处理机制
	defer func() {
		// 捕获当前函数出现的异常，并恢复程序继续执行。
		if x := recover(); x != nil {
			fmt.Println(x)
			b = true
		}
	}()
	f()
	return
}
func main() {
	throwsPanic(hasPanic)
}
