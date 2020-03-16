package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		select {
		case <- time.NewTimer(time.Second).C:
			fmt.Println("timer")
		}
	}
}
