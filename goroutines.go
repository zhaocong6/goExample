package main

import (
	"fmt"
	"sync"
)

func main() {
	//go协程除了有高并发能力
	//go协程还可以调度cpu, 相当于一个轻量级的线程
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		wg.Done()
		fmt.Println(Calculation(1, 2))
	}()

	go func() {
		wg.Done()
		fmt.Println(Calculation(2, 3))
	}()

	//切记使用sleep校验协程是否执行成功是不可靠的
	//应该使用channel, 或者go自带的wg channel包实现
	wg.Wait()
}

func Calculation(a, b int) int {
	return a + b
}
