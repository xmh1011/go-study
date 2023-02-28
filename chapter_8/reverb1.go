package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// 并发回声服务器
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
		go handleConnReverb(conn) // 并发处理客户端的连接请求
	}
}

// handleConn 函数接收一个 net.Conn 类型的参数
// net.Conn 类型的变量，用于表示客户端的连接
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConnReverb(c net.Conn) {
	input := bufio.NewScanner(c) // 创建一个新的Scanner，从c中读取数据
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	// 注意：忽略input.Err()中可能的错误
	c.Close()
}
