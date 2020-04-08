package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGHUP)

	sig := <-ch

	fmt.Println("exit signal:", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	<-ctx.Done()
	log.Println("shutting down")
}
