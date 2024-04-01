package main

import "fmt"

// 指针： * 与 取址 &
// 指针类型的值永远是内存地址
func main() {
	var a int = 4
	var b rune
	var c float32
	var ptr *int // 指针类型int

	fmt.Printf("a type is: %T\n", a)
	fmt.Printf("b type is: %T\n", b)
	fmt.Printf("c type is: %T\n", c)

	// * 与 &
	ptr = &a // 将a的内存地址赋值给int指针类型 ptr
	fmt.Printf("ptr type is: %T\n", ptr)
	fmt.Printf("a value is: %d\n", a)
	fmt.Printf("ptr value is: %p\n", ptr)   // ptr的值
	fmt.Printf("*ptr value is: %d\n", *ptr) // ptr内存空间存储的值
}
