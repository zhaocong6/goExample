package main

import (
	"log"
	"os"
)

func init() {
	//初始化增加时间和文件行数
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	f := NewFile()
	defer f.Close()

	log.SetOutput(f)
	log.Flags()
	debugLog := log.New(f, "[Debug]", log.LstdFlags|log.Lshortfile)
	infoLog := log.New(f, "[Info]", log.LstdFlags|log.Lshortfile)

	defer func() {
		if err := recover(); err != nil {
			debugLog.Println(err)
			infoLog.Println(err)
			log.Println(err)
		}
	}()

	panic("测试log")
}

func NewFile() *os.File {
	f, err := os.OpenFile("./test.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Panic(err)
	}

	return f
}
