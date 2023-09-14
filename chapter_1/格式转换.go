package main

import (
	"fmt"
	"math"
)

func main() {
	// 定义一个[]byte类型的数据
	b := []byte{'a', 'b', 'c'}

	// []byte类型转化为string类型
	s := string(b)

	// 输出结果验证是否转换成功
	fmt.Println("[]byte:", b)
	fmt.Println("string:", s)
	fmt.Println("number:", math.MaxInt64)
	var a int64
	a = 9223372036854775809
	fmt.Println("a:", a)
}
