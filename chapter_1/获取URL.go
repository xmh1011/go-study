package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] { //os.Args[1:]表示从第二个参数开始，也就是从第一个url开始
		resp, err := http.Get(url) // http.Get(url) 会返回一个 http.Response 和一个 error
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body) // io.Copy(dst, src) 会从 src 中读取内容，并将读到的结果写入到 dst 中
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1) // os.Exit(1) 表示非正常退出
		}
		fmt.Printf("%s", b)
	}
}
