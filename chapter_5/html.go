package main

import (
	"fmt"
	"html"
	"io"
	"os"
)

// Node 是一个 HTML 元素或文本节点。
type Node struct {
	Type                    NodeType
	Data                    string      // 元素的标签名或文本
	Attr                    []Attribute // 元素的属性
	FirstChild, NextSibling *Node       // 第一个子节点和下一个兄弟节点
}

type NodeType int32

const ( // 定义节点类型
	ErrorNode    NodeType = iota // 错误
	TextNode                     // 文本节点
	DocumentNode                 // 文档节点
	ElementNode                  // 元素节点
	CommentNode                  // 注释节点
	DoctypeNode                  // 文档类型声明节点
)

// Attribute是一个HTML属性，包含属性名和属性值
type Attribute struct {
	Key, Val string
}

// Parse 从 r 中读取 HTML 并解析，返回解析后的文档树。
func Parse(r io.Reader) (*Node, error) {
	doc, err := Parse(r)
	if err != nil {
		return nil, err
	}
	return forEachNode(nil, doc), nil // 递归遍历文档树
}

// forEachNode 对每个节点 x 调用 pre(x) 和 post(x)。
func forEachNode(t interface{}, doc *Node) *Node {
	if doc.Type == ElementNode {
		t = startElement(t, doc.Data, doc.Attr) // 遇到元素节点时调用 startElement
	}
	if doc.Type == TextNode {
		t = endElement(t, doc.Data) // 遇到文本节点时调用 endElement
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		t = forEachNode(t, c) // 递归遍历子节点
	}
	if doc.Type == ElementNode {
		t = endElement(t, doc.Data) // 遇到元素节点时调用 endElement
	}
	return doc // 返回文档树
}

// endElement 在遇到文本节点时调用。
func endElement(t interface{}, data string) interface{} {
	return t
}

// startElement 在遇到元素节点时调用。
func startElement(t interface{}, data string, attr []Attribute) interface{} {
	if data == "script" || data == "style" { // 跳过脚本和样式元素
		return nil
	}
	for _, a := range attr { // 打印元素的属性
		if a.Key == "href" || a.Key == "src" { // 打印链接和脚本
			fmt.Printf("%s %s", data, html.EscapeString(a.Val)) // html.EscapeString 对字符串进行转义
		}
	}
	return t
}

func main() {
	doc, err := Parse(os.Stdin) // 从标准输入读取 HTML
	if err != nil {             // 解析失败
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err) // 打印错误信息
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) { // 遍历文档树
		fmt.Println(link) // 打印链接
	}
}

// visit 将每个链接添加到 links 中，并返回结果。
func visit(links []string, n *Node) []string {
	if n.Type == ElementNode && n.Data == "a" {
		for _, a := range n.Attr { // 遍历属性
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

