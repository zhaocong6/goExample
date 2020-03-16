package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	_, err := curlGoogle()

	if err != nil {
		fmt.Println(err)
	}
}

func curlGoogle() (bool, error) {
	//假设访问一个不存在的网站
	go func() {
		fmt.Println("curl google ing")
		time.Sleep(time.Second * 60)
	}()

	//设置一个1秒超时
	timer := time.NewTicker(time.Second)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			return false, errors.New("timout")
		}
	}
}
