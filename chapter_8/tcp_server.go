package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var TcpAddr = "localhost"
var TcpPort = 8000
var Length = 108

// func main() {
// 	var FilePath = "/Users/xiaominghao/code/radar/"
// 	var FileName = "radar.txt"
//
// 	filename := FilePath + FileName
// 	data := ReadDataFromFile(filename)
// 	// fmt.Println("1.data: ", data)
// 	err := Server(TcpAddr, TcpPort, data)
// 	fmt.Println(1111)
// 	if err != nil {
// 		fmt.Errorf("Error: %s", err)
// 	}
// }

func main() {
	var FilePath = "/Users/xiaominghao/code/radar/"
	var FileName = "radar.txt"
	
	filename := FilePath + FileName
	data := ReadDataFromFile(filename) // 固定端口
	// _, err := net.Listen("tcp", TcpAddr+":"+strconv.Itoa(TcpPort))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	
	for {
		conn, err := net.Dial("tcp", TcpAddr+":"+strconv.Itoa(TcpPort))
		if err != nil {
			fmt.Println("Error connecting:", err)
			time.Sleep(time.Second) // 等待1秒后重试连接
			continue
		}
		
		fmt.Println("Connected to server:", conn.RemoteAddr())
		
		// 将[]string转换为[]byte
		temp := strings.Join(data, "\n")
		for {
			_, err = conn.Write([]byte(temp)) // 发送数据
			fmt.Println(temp)
			fmt.Println("发送数据成功")
			if err != nil {
				fmt.Println("Error sending data:", err)
				break
			}
			
			time.Sleep(time.Second) // 每秒发送一次数据
		}
		
		conn.Close() // 关闭连接
	}
}

// 从txt文件中读取数据
func ReadDataFromFile(filename string) (data []string) {
	var result string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		result = scanner.Text()
	}
	// fmt.Println("result: ", result)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// // 将result中的数据按照空格分割
	// temp := strings.Split(result, " ")
	// // 将temp中的数据每54位分割一次，存储到data中
	// fmt.Println("temp: ", temp)
	// fmt.Println("len(temp): ", len(temp))
	// 将result中的空格删除，变成新的字符串
	result = strings.Replace(result, " ", "", -1)
	// 将result中的数据，每108位分割一次，存储到data中
	for i := 0; i < len(result); i += Length {
		// data = append(data, strings.Join(result[i:i+Length], ""))
		data = append(data, result[i:i+Length])
	}
	return data
}

// 向固定端口发送数据
func Server(TcpAddr string, TcpPort int, data []string) error {
	for {
		fmt.Println("Server")
		_, err := net.Listen("tcp", TcpAddr+":"+strconv.Itoa(TcpPort))
		fmt.Println(TcpAddr + ":" + strconv.Itoa(TcpPort)) // 创建一个监听器，监听端口 8000
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("server循环")
		// conn 是一个 net.Conn 类型的变量，用于表示客户端的连接
		// conn, err := listener.Accept() // 接收客户端的连接请求
		// if err != nil {
		// 	fmt.Println("server循环失败")
		// 	// 如果接收客户端的连接请求失败，则打印错误信息，然后继续接收下一个客户端的连接请求
		// 	fmt.Println(err)
		// }
		// // 将[]string转换为string
		//
		// temp := strings.Join(data, "\n")
		// fmt.Println("temp:", temp)
		// go func() {
		// 	_, err := io.WriteString(conn, temp) // 向客户端发送当前时间
		// 	if err != nil {
		// 		return
		// 	}
		// }() // 并发处理客户端的连接请求
		// // 在前面添加一个go关键字，使函数在自己的goroutine中执行，而不是在主goroutine中执行
		conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{IP: net.ParseIP(TcpAddr), Port: TcpPort})
		fmt.Println(TcpAddr + ":" + strconv.Itoa(TcpPort))
		if err != nil {
			// fmt.Println("Error connecting:", err)
			return fmt.Errorf("Error connecting: %v", err)
		}
		defer conn.Close()
		
		// 将[]string转换为[]byte
		temp := strings.Join(data, "\n")
		
		// 发送数据
		_, err = conn.Write([]byte(temp))
		if err != nil {
			// fmt.Println("Error sending data:", err)
			return fmt.Errorf("Error sending data: %v", err)
		}
		fmt.Println("Data sent successfully.")
	}
	return nil
}
