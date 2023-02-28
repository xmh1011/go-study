package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000") // 创建一个监听器，监听端口 8000
	if err != nil {
		log.Fatal(err)
	}
	for {
		// conn 是一个 net.Conn 类型的变量，用于表示客户端的连接
		conn, err := listener.Accept() // 接收客户端的连接请求
		if err != nil {
			// 如果接收客户端的连接请求失败，则打印错误信息，然后继续接收下一个客户端的连接请求
			log.Print(err)
			continue
		}
		go handleConn(conn) // 并发处理客户端的连接请求
		// 在前面添加一个go关键字，使函数在自己的goroutine中执行，而不是在主goroutine中执行
	}
}
