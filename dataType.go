package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	//字符串连接
	fmt.Println("go" + "lang")

	fmt.Println(strings.Join([]string{"go", "lang"}, ""))

	var buf bytes.Buffer
	buf.Write([]byte("go"))
	buf.Write([]byte("lang"))
	fmt.Println(buf.String())

	var buf2 bytes.Buffer
	byteStr := []byte("golang")
	for _, b := range byteStr{
		buf2.WriteByte(b)
	}
	fmt.Println(buf2.String())

	//整数和浮点数
	fmt.Println(1+1)
	fmt.Println(1.0/1.0)

	// 布尔型，以及常见的布尔操作。
	fmt.Println(true)
	fmt.Println(false)
	fmt.Println(!true)
	fmt.Println(true == false)
}
