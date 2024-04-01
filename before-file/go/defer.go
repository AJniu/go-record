package main

import "fmt"

// defer关键字
// 作用：延迟执行关键字后的代码到函数return之前执行
// 特点：defer 遵循 LIFO(先进后出)原则
// 良好习惯： 最好将需要defer的代码写在函数顶部，然后写其他正常代码

func practiceDefer() {
	defer fmt.Println("编写在第一条的defer")
	defer fmt.Println("编写在第二条的defer")
	defer fmt.Println("编写在第三条的defer")

	fmt.Println("defer 之后第一条")
	fmt.Println("defer 之后第二条")
	return
}
func main() {
	practiceDefer()
}
