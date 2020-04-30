package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Mutex
	m.Lock()
	c := sync.NewCond(&m)

	go func() {
		m.Lock()
		defer m.Unlock()
		fmt.Println("3. goroutine is owner of lock")
		time.Sleep(2 * time.Second)
		c.Broadcast() //唤醒所有等待的 Wait
		fmt.Println("4. goroutine will release lock soon (deffered Unlock)")
	}()

	fmt.Println("1. main goroutine is owner of lock")
	time.Sleep(1 * time.Second)
	fmt.Println("2. main goroutine is still lockek")
	c.Wait()
	m.Unlock()

	m.Lock()
	go func() {
		m.Lock()
		defer m.Unlock()
		fmt.Println("5. goroutine is owner of lock")
		time.Sleep(2 * time.Second)
		c.Broadcast() //唤醒所有等待的 Wait
		fmt.Println("6. goroutine will release lock soon (deffered Unlock)")
	}()
	c.Wait()
	fmt.Println("Done")
	m.Unlock()
}
