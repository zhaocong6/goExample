package main

import (
	"errors"
	"fmt"
)

type argError struct {
	line string
	code int
}

func (e argError) Error() string {
	return fmt.Sprintf("%d - %s", e.code, e.line)
}

func main() {
	err := f1(0)
	if err != nil {
		fmt.Println(err)
	}

	err = f1(1)
	if err != nil {
		fmt.Println(err)
	}

	err = f2(0)
	if err != nil {
		fmt.Println(err)
	}

	err = f2(1)
	if _, ok := err.(argError); ok {
		fmt.Println("参数错误")
		fmt.Println(err)
		//转换成字符串
		fmt.Println(err.Error())
	}
}

//返回一般的错误
func f1(num int) error {
	if num == 1 {
		return errors.New("不能等于1")
	}

	return nil
}

//返回错误结构体
func f2(num int) error {
	if num == 1 {
		return argError{code: 1, line: "41"}
	}

	return nil
}
