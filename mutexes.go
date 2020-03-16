package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var params struct {
	ops uint64
	mu  sync.Mutex
}

func main() {
	for i := 0; i < 50; i++ {
		go func() {
			for {
				//原子增加
				func() {
					defer params.mu.Unlock()
					params.mu.Lock()
					params.ops++
				}()
				//让出执行时间片
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	params.mu.Lock()
	fmt.Println(params.ops)
	params.mu.Unlock()
}
