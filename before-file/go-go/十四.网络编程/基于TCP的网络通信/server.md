#### 服务端

```go
package main

import (
	"fmt"
	"net"
)

func main() {
	// 服务端代码
	fmt.Println("server is begining！")

	// 使用net包下的Listen()方法监听某端口的请求
	// listen()方法接收两个参数， protocol（协议）和 address（地址）
	// 放回网络监听器接口Listener和错误信息error
	listener, err := net.Listen("tcp", "127.0.0.1:8888")

	if err != nil {
		fmt.Println("服务端监听失败：", err)
		return
	}

	// 如果监听成功之后
	// 循环等待客户端链接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("客户端链接出错：", err)
		} else {
			fmt.Printf("客户端连接成功，con = %v, 接收到的客户端信息为：%v \n", conn, conn.RemoteAddr().String())
		}
	}
}

```

