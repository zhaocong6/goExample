package main

import (
	pd "goExample/streamGrpc/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"math/rand"
	"net"
	"time"
)

type StreamService struct{}

//服务端流
func (s *StreamService) List(r *pd.StreamRequest, stream pd.StreamService_ListServer) error {
	for n := 0; n <= 10; n++ {
		err := stream.Send(&pd.StreamResponse{
			Pt: &pd.StreamPoint{
				Name:  r.Pt.Name,
				Value: int64(n),
			},
		})

		if err != nil {
			return nil
		}
	}

	return nil
}

func (s *StreamService) Record(stream pd.StreamService_RecordServer) error {
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&pd.StreamResponse{
				Pt: &pd.StreamPoint{
					Name:  "客户端流, 响应成功",
					Value: 0,
				},
			})
			return nil
		}

		if err != nil {
			return err
		}

		log.Println(r.Pt.Name, r.Pt.Value)
	}

	return nil
}

func (s *StreamService) Route(stream pd.StreamService_RouteServer) error {

	go func() {
		for {
			stream.Send(&pd.StreamResponse{
				Pt: &pd.StreamPoint{
					Name:  "服务端",
					Value: int64(rand.Int()),
				},
			})

			time.Sleep(time.Millisecond * 100)
		}
	}()

	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		log.Println(r.Pt.Name, r.Pt.Value)
	}
}

func main() {
	server := grpc.NewServer()
	pd.RegisterStreamServiceServer(server, &StreamService{})

	lis, err := net.Listen("tcp", ":9002")
	if err != nil {
		log.Panicln(err)
	}

	err = server.Serve(lis)
	if err != nil {
		log.Panicln(err)
	}
}
