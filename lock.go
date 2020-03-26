package main

import (
	"fmt"
	"github.com/zhaocong6/goUtils/chanlock"
	"time"
)

func main() {
	var lock chanlock.ChanLock

	if lock.TryLock(time.Millisecond) {
		go func() {
			defer lock.Unlock()
			fmt.Println(1)
			time.Sleep(time.Millisecond)
		}()
	}

	if lock.TryLock(time.Millisecond) == false {
		fmt.Println("false")
	} else {
		fmt.Println("true")
	}

	time.Sleep(time.Second)
}
