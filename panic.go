package main

import "fmt"

func main() {
	tryPanic()
	tryPanic()
}

func tryPanic() {

	defer func() {
		//处理panic, 上一级的语句继续执行
		if err := recover(); err != nil {
			fmt.Println(err, "被拦截")
		}
	}()

	panic("panic")
}
