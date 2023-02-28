package main

import (
	"io"
	"log"
	"net"
	"time"
)

// 如何使用该程序？
// 1. 在命令行中，进入本程序所在的目录
// 2. 执行 go build clock1.go 命令，生成 clock1 程序
// 3. 执行 ./clock1 命令，启动 clock1 程序
// 4. 打开另一个命令行窗口，执行 nc localhost 8000 命令，连接到 clock1 程序

// 本程序的功能是，每隔一秒，向客户端发送当前时间
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
		handleConn(conn) // 并发处理客户端的连接请求
	}
}

// handleConn 函数接收一个 net.Conn 类型的参数，然后向客户端发送当前时间
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n")) // 向客户端发送当前时间
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second) // 每隔一秒，向客户端发送当前时间
	}
}
