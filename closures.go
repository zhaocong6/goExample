package main

import "fmt"

func main() {
	test := test()
	fmt.Println(test())

	test2 := func() string {
		return "我也是一个闭包"
	}

	fmt.Println(test2())
}

func test() func() string {
	return func() string {
		return "我是一个闭包"
	}
}
