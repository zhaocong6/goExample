package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	filter := bson.D{{"name", "小红"}}

	//查询一个
	var s *Student
	err = collection.FindOne(context.TODO(), filter).Decode(&s)
	if err == nil {
		fmt.Println(s)
	}

	//游标查询
	op := options.Find()
	op.SetLimit(10)

	var res []*Student

	l, err := collection.Find(context.TODO(), bson.D{{}}, op)
	if err != nil {
		log.Fatalln(err)
	}

	for l.Next(context.TODO()) {
		var elem *Student
		err = l.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		res = append(res, elem)
	}

	l.Close(context.TODO())
	fmt.Println(res)
}
