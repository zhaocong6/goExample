package main

import (
	"flag"
	"fmt"
	"time"
)

var t = flag.Duration("time", 1*time.Second, "Duration")

func main() {
	flag.Parse()

	//go run demo-3.go -time 1s
	//go run demo-3.go -time 1m
	//...
	fmt.Println(t)
}
