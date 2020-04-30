package main

import (
	"fmt"
	json "github.com/json-iterator/go"
	"time"
)

type data struct {
	Name     string `json:"name,omitempty"` //使用omitempty 当该该字段为零值时, json中不会出现
	Age      int    `json:"age"`
	CreateAt string
}

func main() {
	d := &data{
		Age:      0,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	j, _ := json.Marshal(d)

	fmt.Println(string(j))
}
