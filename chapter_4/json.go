package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type movie struct {
	Title  string
	Year   int  `json:"released"`        // 通过tag来指定json中的key
	Color  bool `json:"color,omitempty"` // omitempty表示如果该字段为空，则不输出到json中
	Actors []string
}

var movies = []movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}

func main() {
	data, err := json.Marshal(movies) // 将movies转换成json格式
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err) // log.Fatalf 会打印错误信息，并退出程序
	}
	fmt.Printf("%s", data)
}
