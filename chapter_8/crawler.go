package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

func main() {
	worklist := make(chan []string)  // 可能有重复的URL列表
	unseenLinks := make(chan string) // 去重
	
	// 向worklist发送命令行参数中的URL
	go func() { worklist <- os.Args[1:] }()
	
	// 创建20个爬虫goroutine来获取每个不可见链接
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}
	
	// 主goroutine对URL进行去重
	// 并把没有被爬取过的URL发送给爬虫goroutine
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

// 令牌是一个计数信号量
// 确保并发请求的数量不会超过20
var tokens = make(chan struct{}, 20)

// Extract 函数向指定的URL发送HTTP GET请求
// 解析返回的HTML页面，并返回页面中的所有链接
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	
	if resp.StatusCode != http.StatusOK {
		err := resp.Body.Close()
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // 忽略不合法的URL
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

// forEachNode 调用pre(x)和post(x)遍历以n为根的树中的每个节点
// 两个函数都是可选的
// pre在子节点被访问前调用
// post在子节点被访问后调用
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	
	if post != nil {
		post(n)
	}
}

// var depth int
//
// func startElement(n *html.Node) {
// 	if n.Type == html.ElementNode {
// 		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
// 		depth++
// 	}
// }
//
// func endElement(n *html.Node) {
// 	if n.Type == html.ElementNode {
// 		depth--
// 		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
// 	}
// }

// Crawl the web breadth-first,
// starting from the command-line arguments.
func crawl(url string) []string {
	// ...
	fmt.Println(url)
	tokens <- struct{}{} // 获取令牌
	list, err := Extract(url)
	<-tokens // 释放令牌
	
	if err != nil {
		log.Print(err)
	}
	return list
}
