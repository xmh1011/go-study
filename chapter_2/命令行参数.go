package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()                               // 解析命令行参数
	fmt.Print(strings.Join(flag.Args(), *sep)) // flag.Args() returns the non-flag command-line arguments
	if !*n {
		fmt.Println()
	}
}
