package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	cl := make(chan interface{})
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()

			for {
				select {
				case <-cl:
					return
				default:
					time.Sleep(time.Second)
					fmt.Println("runing")
				}
			}
		}()
	}

	go func() {
		time.Sleep(time.Millisecond * 50)
		fmt.Println("send end")
		close(cl)

		//防止部分程序假死, 须停止服务
		select {
		case <-time.After(time.Second * 60):
			os.Exit(0)
		}
	}()

	wg.Wait()
	fmt.Println("end")
}
