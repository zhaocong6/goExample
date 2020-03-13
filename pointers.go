package main

import "fmt"

func main() {
	i := 0

	zeroval(i)
	fmt.Println(i)

	zeroptr(&i)
	fmt.Println(i)
}

//ival是一个内存拷贝
func zeroval(ival int) {
	//修改了局部变量
	ival = 1
}

//iptr是指针的拷贝
func zeroptr(iptr *int) {
	//修改指针的值
	*iptr = 1

	//这个修改无效, iptr是一个拷贝, 局部的变量
	i := 100
	iptr = &i
}