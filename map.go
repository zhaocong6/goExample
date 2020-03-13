package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[1] = 1
	m[3] = 3
	m[2] = 2

	//如果使用new 就需要初始化内存, new默认是零值
	m2 := new(map[int]int)
	*m2 = map[int]int{}
	(*m2)[1] = 1

	fmt.Println(m, m2)
}
