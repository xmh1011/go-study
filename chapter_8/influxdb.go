package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
	url := "http://localhost:8086/query"
	query := "CREATE DATABASE " + "db1"
	data := []byte(fmt.Sprintf("q=%s", query))
	
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error creating request: %v", err)
	}
	
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error making request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("openGemini create database error: %s", resp.Status)
	}
	for {
		data := "test,location=server1 value1=0.64,value2=0.12 " + strconv.FormatInt(time.Now().UnixNano(), 10)
		// 构造请求
		req, err := http.NewRequest("POST", "http://localhost:8086/write?db=db1", bytes.NewBufferString(data))
		
		// 执行curl
		if err != nil {
			fmt.Println("Invalid request: %s", err)
		}
		req.Header.Set("Content-Type", "application/octet-stream")
		
		// 发送请求
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making request: %s", err)
		}
		
		// 检查响应状态码
		if resp.StatusCode != http.StatusOK {
			fmt.Println("openGemini insert data error: %s", resp.Status, data)
		}
		
		time.Sleep(10 * time.Second)
	}
}
