package main

import "fmt"

type job chan int
type result chan string

func main() {
	//简单的使用goroutine实现work 没有实现安全退出, 真实环境不可用
	jobs := make(job, 100)
	result := make(result)

	//创建3个work
	for i := 1; i <= 3; i++ {
		go work(i, jobs, result)
	}

	//抛出9个任务
	for i := 1; i <= 9; i++ {
		i := i
		go func() {
			jobs <- i
		}()
	}

	//等待指定任务运行完成
	for i := 0; i < 9; i++ {
		<-result
	}
}

func work(workerId int, j job, res result) {
	for no := range j {
		fmt.Println("worker", workerId, "processing job", no)
		res <- "success " + string(no)
	}
}
