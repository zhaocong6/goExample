package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	_, err := timout()

	if err != nil {
		fmt.Println(err)
	}
}

func timout() (bool, error) {
	timer := time.NewTicker(time.Second)

	go func() {
		fmt.Println("curl google ing")
		time.Sleep(time.Second * 60)
	}()

	for {
		select {
		case <-timer.C:
			timer.Stop()
			return false, errors.New("timout")
		}
	}
}
