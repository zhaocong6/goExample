package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type data struct {
	//tag 美化输出
	Name     string `json:"name"`
	Age      int    `json:"age"`
	CreateAt string
}

func main() {
	//注意: 必须是可以导出的字段
	d := &data{
		Name:     "测试",
		Age:      0,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	j, _ := json.Marshal(d)

	fmt.Println(string(j))
}
