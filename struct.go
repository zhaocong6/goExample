package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	//初始化一个结构体
	fmt.Println(person{"Bob", 20})
	fmt.Println(&person{"Bob", 20})
	person := &person{"Bob", 20}
	fmt.Println(person)
	fmt.Println(person.name)
}
