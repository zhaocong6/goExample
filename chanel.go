package main

import "fmt"

func main() {
	var ch chan int

	//创建一个无缓存的channel
	//注意: 这里的顺序必须是先准备读取chan, 再写入chan, 否则会死锁
	ch = make(chan int)
	go func() {
		defer func() { ch <- 1 }()
	}()
	fmt.Printf("无缓存ch%d\r\n", <-ch)

	//创建一个有缓存的延迟channel
	ch = make(chan int, 2)
	go func() {
		defer func() { ch <- 2 }()
	}()
	fmt.Printf("延迟ch%d\r\n", <-ch)

	//创建一个有缓存channel
	ch = make(chan int, 3)
	go func() {
		defer func() { ch <- 3 }()
	}()
	fmt.Printf("缓存ch%d\r\n", <-ch)

	//使用chan进行通讯
	pingCh := make(chan string)
	pongCh := make(chan string)
	go ping(pingCh)
	go pong(pingCh, pongCh)

	fmt.Println(<-pongCh)
}

func ping(pingCh chan string) {
	pingCh <- "ping"
}

func pong(pingCh chan string, pongCh chan string) {
	ping := <-pingCh
	if ping == "ping" {
		pongCh <- "pong"
	}
}
