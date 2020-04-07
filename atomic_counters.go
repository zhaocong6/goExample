package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			for {
				//原子增加
				atomic.AddUint64(&ops, 1)

				//让出执行时间片
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	//原子获取
	fmt.Println(atomic.LoadUint64(&ops))
}
