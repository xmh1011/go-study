package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) // 创建一个 channel
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine fetch
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds()) // time.Since(start) 返回一个 time.Duration 类型的值
}
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	// io.Copy(dst, src) 会从 src 中读取内容，并将读到的结果写入到 dst 中
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // ioutil.Discard 是一个 io.Writer，它会忽略所有写入的数据
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err) // send to channel ch
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
