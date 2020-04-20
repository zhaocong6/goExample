package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os/user"
)

func main() {
	//使用os/user包 不能交叉编译
	u, _ := user.Current()
	fmt.Println(u.HomeDir)

	//使用第三方
	d, _ := homedir.Dir()
	fmt.Println(d)
}
