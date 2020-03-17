package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main () {
	//strings包简单使用

	//是否包含, 返回bool
	p(s.Contains("test", "e"))

	//计算出现次数, 返回int
	p(s.Count("test", "t"))

	//获取是否有这个前缀
	p(s.HasPrefix("test", "te"))

	//判断是否有这个后缀
	p(s.HasSuffix("test", "st"))

	//判断字符索引位置(首次出现)
	p(s.Index("test", "t"))

	//将一个字符切片合并成一个字符串
	p(s.Join([]string{"te", "st"}, ""))

	//重复
	p(s.Repeat("e", 5))

	//替换
	p(s.Replace("test", "t", "e", 1))

	//将一个字符串分割成一个字符切片
	p(s.Split("t-e-s-t", "-"))
}
