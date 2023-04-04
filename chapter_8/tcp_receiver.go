package main

import (
	"fmt"
	"net"
)

func main() {
	ip := "127.0.0.1" // 指定IP地址
	port := "8000"    // 指定端口号
	
	// 建立TCP连接
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()
	
	// 循环接收数据
	for {
		// 读取数据
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		
		// 处理数据
		fmt.Println(string(data[:n]))
	}
}
