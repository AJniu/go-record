## fmt

fmt 包实现了类似 C语言 printf 和 scanf 的格式化 I/O 。

主要分为 向外输出内容 和 获取输入内容 两大部分。



##### 向外输出：标准库 fmt 提供了以下几种输出相关函数。

##### 1. Print

​	Print 系列函数会将内容输出到系统的标准输出：

1. Print 函数直接输出内容（不换行）。
2. Printf 函数支持格式化输出字符串（不换行）。
3. Println 函数会在输出内容的结尾添加一个换行符。

```go
// any 即 interface{} 别名
// Print系列函数都会返回输出的字符串长度 和 错误
func Print(a ...any) (n int, err error) 

func Printf(format string, ...any) (n int, err error)

func Println(a ...any) (n int, err error)
```

##### 2. Fprint

​	Fprint 系列函数会将内容输出到一个 io.Writer 接口类型的变量 w 中，通常用这个函数往文件中写入内容。

​	1.
