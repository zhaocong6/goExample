package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	//注意 不使用break, case也不会自动继续执行, 使用break主要是为了控制一些流程, 早点结束switch
	switch i {
	case 1:
		fmt.Println("i等于1")
	case 2:
		fmt.Println("i等于2")
	case 3:
		fmt.Println("i等于3")
	}

	switch {
	case time.Now().Hour() == 1:
		fmt.Println("当前是1点")
	}

	what := func(v interface{}) {
		switch v.(type) {
		case int:
			fmt.Println("int类型")
		case bool:
			fmt.Println("bool类型")
		default:
			fmt.Println("其他类型")
		}
	}

	what(1)
	what(true)
	what("1")
}
