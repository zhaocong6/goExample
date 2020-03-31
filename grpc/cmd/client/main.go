package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pd "github.com/zhaocong6/proto/hello"
	"log"
)

func main() {
	Request()
}

func Request() {
	conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	c := pd.NewHelloClient(conn)
	req := &pd.HelloRequest{Name:"gRpc"}
	res, err := c.SayHello(context.Background(), req)
	if err != nil {
		grpclog.Fatalln(err)
	}

	log.Println(res.Message)
}
