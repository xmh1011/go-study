package main

import (
	"io"
	"log"
	"net"
	"os"
)

// 如何使用该程序？
// 1. 在命令行中，进入本程序所在的目录
// 2. 执行 go build netcat1.go 命令，生成 netcat1 程序
// 3. 执行 ./netcat1 命令，启动 netcat1 程序
// 4. 打开另一个命令行窗口，执行 nc localhost 8000 命令，连接到 netcat1 程序

func main() {
	// conn 是一个 net.Conn 类型的变量，用于表示客户端的连接
	conn, err := net.Dial("tcp", ":8000") // 创建一个连接，连接到端口 8000
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()        // 关闭连接
	mustCopy(os.Stdout, conn) // 从标准输入中读取数据，然后写入到连接中
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

// mustCopy是一个io.Copy的包装器，如果io.Copy返回一个错误，那么就会调用log.Fatal函数，终止程序的执行
