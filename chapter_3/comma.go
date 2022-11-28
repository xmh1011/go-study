package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(comma("1234567889988"))
	strconv.Itoa(1) // 将int转换成asc码
	a := 5
	b := 4 << a
	fmt.Printf("b:\n")
	fmt.Println(b)

}

func comma(s string) string { // comma inserts commas in a non-negative decimal integer string.
	var newByte byte = ','
	n := len(s)
	buf := bytes.NewBuffer([]byte{}) // bytes.NewBuffer returns a new Buffer initialized with the given bytes.

	if n <= 3 {
		return s
	}

	for i := 0; i < n; i++ {
		if (n-i)%3 == 0 && i != 0 {
			buf.WriteByte(newByte) // WriteByte appends the byte c to the buffer, growing the buffer as needed.
		}
		buf.WriteByte(s[i])
	}

	return buf.String()
}
