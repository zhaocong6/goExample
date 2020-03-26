package main

import "fmt"

type boy struct {
	Name string
	Age  int
}

func main() {
	bob := &boy{
		Name: "bob",
		Age:  12,
	}

	//%v 打印出结构体数据
	fmt.Printf("%v", bob)
	//%+v 打印出结构体数据和字段
	fmt.Printf("%+v", bob)

	//打印出boole
	t := true
	fmt.Printf("%t", t)

	//等等 https://blog.csdn.net/chenbaoke/article/details/39932845
	i := 10086
	fmt.Printf("%b\r\n", i)

	f := 11.123
	fmt.Printf("%.1f\r\n", f)
	fmt.Printf("%.3g\r\n", f)

	s := "test"
	fmt.Printf("%*.*s\r\n", 2, 3, s)
}
