package main

import (
	"fmt"
	"os"
)

func main() {

	ch := make(chan int)
	ch2 := make(chan int)

	go func() {
		fmt.Println("ch1")
		defer close(ch)
	}()

	go func() {
		fmt.Println("ch2")
		defer close(ch2)
	}()

	for {
		//类型linux中的select函数
		select {
		case <-ch:
			os.Exit(0)
		case <-ch2:
			os.Exit(0)
		}
	}
}
