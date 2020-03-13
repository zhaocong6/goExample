package main

import "fmt"

//有参数的接口
type Animal interface {
	Speak() string
}

type Dog struct {}

type Cat map[string]string

func main() {
	animals := []Animal{Dog{}, &Cat{}}
	for _, a := range animals{
		fmt.Println(a.Speak())
	}
}

func (d Dog) Speak() string {
	return "wang wang"
}

func (c *Cat) Speak() string {
	return "miao miao "
}
