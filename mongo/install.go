package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	//创建链接信息
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

	//链接mongo
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Panicln(err)
	}
	defer client.Disconnect(context.TODO())

	//设置一个数据, 选择一个表
	collection := client.Database("test").Collection("study")
	s1 := &Student{
		Name: "小明",
		Age:  10,
	}

	s2 := &Student{
		Name: "小红",
		Age:  10,
	}

	students := []interface{}{s1, s2}

	res1, err := collection.InsertOne(context.TODO(), students[0])
	res2, err := collection.InsertMany(context.TODO(), students)
	fmt.Println(res1, res2)
}
