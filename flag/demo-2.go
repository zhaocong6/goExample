package main

import (
	"flag"
	"fmt"
)

var intFlag int

func init() {
	//短选项定义
	flag.IntVar(&intFlag, "intFlag", 0, "test int flag")
	flag.IntVar(&intFlag, "i", 0, "test int flag(shorthand)")
}

func main() {
	//解析os.Args
	flag.Parse()

	//go run demo-2.go -intFlag 2 -i 4
	fmt.Println(intFlag)
}
