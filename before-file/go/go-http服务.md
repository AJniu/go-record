### go 是如何创建一个服务的？

+ 调用`http`包下的`ListenAndServe()`

```go
http.HandleFunc("/", handleFunc) // 设置访问的路由及对应的处理函数
// 监听端口 9090，并创建后端服务
// 如果一切正常，该函数将一直阻塞，直到服务关闭或发生错误。
err := http.ListenAndServe(":9090", nil)
```



###### `http.ListenAndServe` 的实现

```go
// 作用：监听并提供 HTTP 服务。

// 接收参数：	
// 1. addr: 字符串类型，表示 HTTP 服务监听的地址和端口号
// 2. handler: Interface Handler { ServeHTTP(ResponseWriter, *Requset)}类型，表示处理接收到的 HTTP 请求的处理器。

// 返回:  error接口类型,表示在监听并提供 HTTP 服务时发生的任何错误。如果一切正常，该函数将一直阻塞，直到服务关闭或者发生错误。

func ListenAndServe(addr string, handler Handler) error {
    // 注意：此函数本身并不提供服务，而是创建 Server 结构体，并调用其 ListenAndServe 方法来提供服务。
    
    // 创建 Server 结构体实例
	server := &Server{Addr: addr, Handler: handler}
    // 调用 server.ListenAndServe() 方法并返回其返回值。
    // 因为其提供的HTTP服务，所以此返回值恒为一个non-nil error。
	return server.ListenAndServe()
}
```

###### 真正提供HTTP服务的 `Server`结构体的`ListenAndServe` 方法 详解

```go
// 作用：阻塞当前协程，直到服务停止或出现错误。

func (srv *Server) ListenAndServe() error {
	if srv.shuttingDown() {
		return ErrServerClosed
	}
	addr := srv.Addr
	if addr == "" {
		addr = ":http"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return srv.Serve(ln)
}
```

分片段详解：

```go
// 片段详解
if srv.shuttingDown() {
    return ErrServerClosed  // var ErrServerClosed = errors.New("http: Server closed")
} 
// 作用：srv.shuttingDown() 是一个内部方法，用于检查当前 HTTP 服务器是否正在关闭或已经关闭。如果服务器正在关闭或已经关闭，则 srv.shuttingDown() 返回 true，否则返回 false。

// 具体来说，srv.shuttingDown() 函数检查 srv.inShutdown 字段的值，如果该字段的值不为 0，则表示服务器正在关闭或已经关闭，此时 srv.shuttingDown() 返回 true。否则，如果 srv.inShutdown 字段的值为 0，则表示服务器正常运行中，此时 srv.shuttingDown() 返回 false。

// srv.inShutdown 字段的值会在服务器关闭时被设置为非 nil 的值，以表示服务器正在关闭或已经关闭。

// Server struct结构体 shuttingDown 方法
func (s *Server) shuttingDown() bool {
	
    // Server struct结构体的 inShutdown 属性 atomicBool 为 int32 的别名 
	// inShutdown atomicBool
    return s.inShutdown.isSet()
}

// atomicBool 类型的 isSet()方法
// 返回当前的 atomicBool 实例是否被设置为 true，即该方法会判断 atomicBool 的底层值是否为 1。
// 这个方法的核心是使用了 Go 语言中的原子操作（atomic operation）。

// 在这个方法中，使用了 atomic.LoadInt32 函数来原子地读取 atomicBool 实例的值，然后将其转换为 int32 类型，接着判断这个 int32 类型的值是否为 0。
func (b *atomicBool) isSet() bool { 
    // 其中 (*int32)(b) 就是将 *atomicBool 类型的b 转换为 *int32 类型
    return atomic.LoadInt32((*int32)(b)) != 0 
}


// 上面代码中的 atomic 的 LoadInt32方法
// 用于以原子操作的方式从指定的内存地址中读取一个 int32 类型的值，并将其返回。
//这个函数的参数 addr 是一个指向 int32 类型变量的指针，表示要读取的内存地址。函数返回一个 int32 类型的值，表示从该地址读取的值。
// 在 Go 语言中，多个协程并发地访问同一个变量可能会出现数据竞争的问题，而使用原子操作可以保证在多个协程并发访问同一个变量时，对该变量的读写操作是原子的，从而避免了数据竞争和并发问题。
// 因此，LoadInt32 函数的使用非常方便和安全，可以用于多个协程并发读取同一个 int32 类型的变量的值。需要注意的是，LoadInt32 函数仅仅是读取指定内存地址上的值，而不会对该值进行修改。如果需要修改该值，需要使用 StoreInt32 函数或其他的原子操作函数。
func LoadInt32(addr *int32) (val int32)


```

```go
// 片段：
	// 监听设置的地址
	addr := srv.Addr
// 如果srv.Addr为空字符串，默认使用 ":http"
	if addr == "" {
		addr = ":http"
	}
// 创建一个监听地址为 addr 的 TCP监听器，监听来自客户端的TCP连接请求
// 如果为默认的 ":http"，则表示监听在本地IP地址的所有可用的网络接口上，并监听端口为 80 的HTTP

// `ListenAndServe`调用了`net.Listen("tcp", addr)`，也就是底层用TCP协议搭建了一个服务，最后调用`src.Serve`监控我们设置的端口。
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return srv.Serve(ln)
```

