package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)                           // index 为向 url发送请求时，调用的函数
	log.Fatal(http.ListenAndServe("localhost:8000", nil)) // ListenAndServe 监听端口，如果端口被占用，会报错
}

func index(w http.ResponseWriter, r *http.Request) { // index 函数接收两个参数，第一个是 http.ResponseWriter，第二个是 *http.Request
	fmt.Fprintf(w, "Golang是世界上最好的语言！")
}
