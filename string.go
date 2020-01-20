package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	//+
	//字符串常量拼接, 会产生很多字符串常量, 给gc带来负担
	fmt.Println("hello + world", "hello" + "world")

	//fmt.Sprintf
	//使用fmt拼接, 通过Sprintf格式化方式, 内部使用 []byte 实现, 不会产生临时的字符串常量, 由于形参使用了interface, 可以使用很多类型, 性能也一般
	fmt.Println("hello + world", fmt.Sprintf("%s%s%d", "hello", "world", 1))

	//strings.Join
	//join会先根据字符串数组的内容，计算出一个拼接之后的长度，然后申请对应大小的内存，一个一个字符串填入，在已有一个数组的情况下，这种效率会很高，但是本来没有，去构造这个数据的代价也不小
	hello := "hello"
	world := "world"
	fmt.Println("hello + world", strings.Join([]string{hello, world}, ""))
	//内部假实现
	helloLen := len(hello)
	worldLen := len(world)

	//创建一个结构体
	var b strings.Builder
	//初始化字符串大概的长度
	b.Grow(helloLen + worldLen)
	//写入字符串
	b.WriteString(hello)
	b.WriteString(world)
	//获取字符串
	fmt.Println(b.String())

	//buffer.WriteString
	//和strings.Join不同, 该函数预分配一个空间, 当下次空间不足时, 当前长度*2 如果超过一定值 就不再*2 而是*1.25
	//不会频繁的分配内存 比strings.Join性能高 最高效的拼接方式
	hello = "hello"
	world = "world"
	var buffer bytes.Buffer
	buffer.WriteString(hello)
	buffer.WriteString(world)
	fmt.Println(buffer.String())

	//在已有字符串数组的场合，使用 strings.Join() 能有比较好的性能
	//在一些性能要求较高的场合，尽量使用 buffer.WriteString() 以获得更好的性能
	//性能要求不太高的场合，直接使用运算符，代码更简短清晰，能获得比较好的可读性
	//如果需要拼接的不仅仅是字符串，还有数字之类的其他需求的话，可以考虑 fmt.Sprintf()
}
