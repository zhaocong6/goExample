package main

import (
	"flag"
	"fmt"
)

func main() {
	args := []string{"-intFlag", "12", "-boolFlag=false"}

	//NewFlagSet方法有两个参数，第一个参数是程序名称，输出帮助或出错时会显示该信息。第二个参数是解析出错时如何处理，有几个选项：
	//ContinueOnError：发生错误后继续解析，CommandLine就是使用这个选项；
	//ExitOnError：出错时调用os.Exit(2)退出程序；
	//PanicOnError：出错时产生 panic。
	fs := flag.NewFlagSet("MyFlagSet", flag.ContinueOnError)

	var intFlag int
	var boolFlag bool
	fs.IntVar(&intFlag, "intFlag", 0, "")
	fs.BoolVar(&boolFlag, "boolFlag", true, "")
	fs.Parse(args)

	fmt.Println(intFlag)
	fmt.Println(boolFlag)
}
