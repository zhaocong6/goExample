package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	withHandler()
	timeoutHandler()
	withDeadlineHandler()
	withValue()
}

func withValue() {
	ctx := context.WithValue(context.Background(), "key", "value")
	val := ctx.Value("key").(string)
	fmt.Println(val)
}

func withDeadlineHandler() {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	go doTimeOutStuff(ctx)
	time.Sleep(time.Second * 3)
}

func timeoutHandler() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	go doTimeOutStuff(ctx)
	time.Sleep(time.Second * 3)
}

func withHandler() {
	ctx, cancel := context.WithCancel(context.Background())
	go doStuff(ctx)
	time.Sleep(time.Second * 10)
	cancel()
	time.Sleep(time.Second * 2)
}

func doStuff(ctx context.Context) {
	for {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("done")
			return
		default:
			fmt.Println("work")
		}
	}
}

func doTimeOutStuff(ctx context.Context) {
	for {
		time.Sleep(time.Second)

		if deadline, ok := ctx.Deadline(); ok {
			fmt.Println("deadline set")

			if time.Now().After(deadline) {
				fmt.Println("deadline end")
				return
			}
		}

		select {
		case <-ctx.Done():
			fmt.Println("done")
			return
		default:
			fmt.Println("work")
		}
	}
}
