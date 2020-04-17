package main

import (
	"flag"
	"fmt"
)

var fInt = flag.Int("intFlag", 0, "test int flag")
var fBool = flag.Bool("boolFlag", false, "test bool flag")

func main() {
	//解析os.Args
	flag.Parse()

	fmt.Println(*fInt)
	fmt.Println(*fBool)

	//输出解析失败的参数
	fmt.Println(flag.Args())
}
