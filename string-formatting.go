package main

import "fmt"

func main() {
	type point struct {
		x, y int
	}

	p := point{1, 2}

	//打印结构体的值
	fmt.Printf("%v\r\n", p)
	//打印结构体的值和键
	fmt.Printf("%+v\r\n", p)
	//打印出值的类型
	fmt.Printf("%T\r\n", p)
	//格式化bool值
	fmt.Printf("%t\r\n", false)

	//格式化标准的十进制
	fmt.Printf("%d\r\n", 123)
	//输出二进制
	fmt.Printf("%b\r\n", 123)
	fmt.Printf("%c\r\n", 122)
}
