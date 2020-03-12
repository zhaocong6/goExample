package main

import "fmt"

func main() {

	//i是一个main函数的局部变量, main结束后, i进入gc
	i := 0
	for {
		if i >= 2 {
			break
		}

		fmt.Println("run")
		i++
	}

	//i 在for中 是一个局部变量, for运行完成以后, 进入gc
	for i := 0; i < 2; i++ {
		fmt.Println("run")
	}

	i2 := 0
	for {
		fmt.Println("run")
		if i2 >= 2 {
			//注意使用continue是跳出该次循环, 这里会一直死循环
			continue
		}

		i2++
	}
}
