package main

import "fmt"

func main() {
	fmt.Println(testFunc())
	fmt.Println(TestFunc())
	variadic(1,2,43,543,6435,6)
}

func testFunc() string {
	return "不能导出的方法"
}

func TestFunc() string {
	return "可以导出的方法"
}

func variadic(nums ...int) {
	fmt.Println(nums)
}

