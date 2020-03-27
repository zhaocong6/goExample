package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			select {
			//这种写法没问题
			case <-time.NewTimer(time.Second).C:
				//fmt.Println("timer")
			}
		}
	}()

	//这种写法会造成内存泄露
	go func() {
		//for {
		//	select {
		//	case <-time.NewTimer(time.Second * 5).C:
		//		fmt.Println("5")
		//	default:
		//		fmt.Println("default")
		//	}
		//
		//	fmt.Println("for")
		//}
	}()

	//修改
	go func() {
		t := time.NewTimer(time.Second * 5)
		defer t.Stop()

		for {
			select {
			case <-t.C:
				fmt.Println("5")
			default:
				time.Sleep(time.Second * 10)
				//t.Reset(time.Second * 5)
				fmt.Println("default")
			}

			fmt.Println("for")
		}
	}()

	select {}
}
