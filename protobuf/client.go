package main

import (
	"context"
	pd "goExample/protobuf/first"
	"google.golang.org/grpc"
	"log"
	"sync"
)

func main() {
	//连接rpc服务
	conn, err := grpc.Dial(":9001", grpc.WithInsecure())
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			//创建rpc客户端, 此时连接可以复用
			client := pd.NewSearchServiceClient(conn)
			//调用rpc方法
			response, err := client.Search(context.TODO(), &pd.SearchRequest{
				Request: "第一次请求rpc",
			})
			if err != nil {
				log.Panicln(err)
			}
			log.Println(response.Response)
		}()
	}

	wg.Wait()
}
