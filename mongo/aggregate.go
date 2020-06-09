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

	pipeline := []bson.M{bson.M{"$match": bson.M{"name": "小红"}}, bson.M{"$limit": 3}}

	cur, err := collection.Aggregate(context.TODO(), pipeline, options.Aggregate())

	if cur != nil {
		for cur.Next(context.TODO()) {
			var res Student
			cur.Decode(&res)
			fmt.Println(res)
		}
	}
}
