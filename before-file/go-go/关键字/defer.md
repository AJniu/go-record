#### defer 关键字

1. 作用：（及时释放资源）在函数中，程序员经常需要创建资源，为了在函数执行完毕后，及时的释放资源，go设计者提供了defer关键字 （看完实际效果感觉这说法不对，记效果就行）

    ```go
    package main
    
    import (
    	"fmt"
    )
    
    func main() {
    	fmt.Println(addFn(10, 20))
    }
    
    // 定义闭包函数
    func addFn(n1 int, n2 int) int {
    	var sum int = 0
    	// 在go中，程序遇到defer关键字，不会立即执行defer中的语句，
    	// 而是将defer后的语句压入一个栈中，然后继续执行函数后面的语句
    	defer fmt.Println("n1 = ", n1)
    	defer fmt.Println("n2 = ", n2)
    
    	// 由于栈的特点是 先进后出
    	// 所以在函数执行完毕以后，从栈中取出语句开始执行，按照现金后出的规则执行语句
    	sum += n1 + n2
    	fmt.Println("sum =", sum)
    	return sum
    }
    
    ```

    

2. 特点：程序执行中，遇到defer关键字，会将后面的代码语句压入栈中，也会将相关的值同时拷贝入栈中，不会随着函数后面的变化而变化。（验证，在上述代码defer后添加 对n1和n2的更改，看n1和n2的打印结果）

