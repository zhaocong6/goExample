package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"sync"
	"time"
)

func main() {
	addr := []string{"localhost:2181"}
	conn, _, err := zk.Connect(addr, time.Second*1)
	if err != nil {
		fmt.Println(err)
	}

	start := time.Now().UnixNano()
	var wg sync.WaitGroup
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			lock := zk.NewLock(conn, "/lock/product-1", zk.WorldACL(zk.PermAll))
			lock.Lock()
			lock.Unlock()
		}()
	}

	wg.Wait()
	end := time.Now().UnixNano()

	fmt.Println((end - start) / 1e6)
}
