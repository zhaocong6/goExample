package main

import (
	"context"
	"fmt"
	pd "github.com/zhaocong6/proto/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

type helloService struct{}

var HelloService = helloService{}

func (h helloService) SayHello(ctx context.Context, in *pd.HelloRequest) (*pd.HelloResponse, error) {
	resp := new(pd.HelloResponse)
	resp.Message = fmt.Sprintf("hello")

	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:50052")
	if err != nil {
		grpclog.Fatalf("服务启动失败!")
	}

	s := grpc.NewServer()

	pd.RegisterHelloServer(s, HelloService)

	fmt.Println("启动成功")

	s.Serve(listen)
}
