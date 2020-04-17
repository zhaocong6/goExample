package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type interval []int

func (i *interval) String() string {
	return fmt.Sprintln(*i)
}

func (i *interval) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}

	//使用规定的, 进行字符串的切割
	for _, v := range strings.Split(value, ",") {
		if toi, err := strconv.Atoi(v); err == nil {
			*i = append(*i, toi)
		}
	}

	return nil
}

var i interval

func init() {
	flag.Var(&i, "flagInt", "test int")
}

func main() {
	flag.Parse()
	//自定义flag, 需要实现flag value接口
	fmt.Println(i)
}
