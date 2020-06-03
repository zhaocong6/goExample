package main

import (
	"context"
	pd "goExample/protobuf/first"
	"google.golang.org/grpc"
	"log"
	"net"
)

type ServiceService struct {
}

//定义一个处理方法
func (s *ServiceService) Search(ctx context.Context, r *pd.SearchRequest) (*pd.SearchResponse, error) {
	return &pd.SearchResponse{
		Response: "测试grpc",
	}, nil
}

func main() {
	server := grpc.NewServer()
	pd.RegisterSearchServiceServer(server, &ServiceService{})

	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Panicln(err)
	}

	err = server.Serve(lis)
	if err != nil {
		log.Panicln(err)
	}
}
