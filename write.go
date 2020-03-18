package main

import (
	"bufio"
	"os"
)

func main() {
	f, _ := os.Create("./text.txt")
	f.Write([]byte("写入数据"))
	//将缓存数据写入磁盘
	f.Sync()

	//使用bufio
	b := bufio.NewWriter(f)
	b.WriteString("测试写入")
	b.Flush()
}
