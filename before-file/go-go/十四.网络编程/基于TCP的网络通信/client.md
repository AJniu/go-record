#### 客户端

1. go网络编程主要使用net包下的内容
2. 客户端与服务端建立连接使用net.Dial(protocol, address string) (Conn, error)方法

```go
package main

import (
	"fmt"
	"net" // net包负责网络编程的相关内容
)

func main() {
	fmt.Println("客户端开始启动")

	// Dial函数主要负责和服务端建立连接
	// Dial函数接收两个参数protocol（协议），地址（address string)
	// Dial函数返回两个值Conn（网络连接），err（错误）
	conn, err := net.Dial("tcp", "127.0.0.1:8888")

	if err != nil {
		fmt.Println("连接服务端失败：", err)
		return
	}
	fmt.Println("连接服务端成功：", conn)
}

```

