package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//go协程除了有高并发能力
	//go协程还可以调度cpu, 相当于一个轻量级的线程
	//go协程 能处理高并发和并行, 与其他语言不同 其他语言一般是单线协程(如 php python)
	//切记使用sleep校验协程是否执行成功是不可靠的
	//应该使用channel, 或者go自带的wg channel包实现

	//使用wg
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println(Calculation(1, 2))
		time.Sleep(time.Second)
	}()

	go func() {
		defer wg.Done()
		fmt.Println(Calculation(2, 3))
		time.Sleep(time.Second)
	}()
	wg.Wait()

	//使用channel
	ch := make(chan int, 2)

	go func() {
		defer func() { ch <- 1 }()
		fmt.Println(Calculation(1, 2))
		time.Sleep(time.Second)
	}()

	go func() {
		defer func() { ch <- 1 }()
		fmt.Println(Calculation(2, 3))
		time.Sleep(time.Second)
	}()

	<-ch
	<-ch
}

func Calculation(a, b int) int {
	return a + b
}
