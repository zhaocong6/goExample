package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	cancel := make(chan bool)

	var wg sync.WaitGroup

	for i := 10; i > 0; i-- {
		wg.Add(1)
		go work(cancel, &wg)
	}
	time.Sleep(time.Millisecond)
	close(cancel)
	wg.Wait()
	fmt.Println("close")
}

func work(cancel chan bool, wg *sync.WaitGroup) {

	defer func() {
		time.Sleep(time.Millisecond * 3000)
		wg.Done()
	}()

	for {
		select {
		default:
			fmt.Println("run")
		case <-cancel:
			return
		}
	}
}
