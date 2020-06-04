package main

import (
	"context"
	"fmt"
	pd "goExample/streamGrpc/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"math/rand"
	"time"
)

func main() {
	conn, err := grpc.Dial(":9002", grpc.WithInsecure())
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	client := pd.NewStreamServiceClient(conn)

	//err = printLists(client, &pd.StreamRequest{Pt: &pd.StreamPoint{Name: "gRPC Stream Client: List", Value: 2018}})
	//if err != nil {
	//	log.Fatalf("printLists.err: %v", err)
	//}

	//err = printRecord(client, &pd.StreamRequest{Pt: &pd.StreamPoint{Name: "gRPC Stream Client: Record", Value: 2018}})
	//if err != nil {
	//	log.Fatalf("printRecord.err: %v", err)
	//}

	err = printRoute(client, &pd.StreamRequest{Pt: &pd.StreamPoint{Name: "gRPC Stream Client: Route", Value: 2018}})
	if err != nil {
		log.Fatalf("printRoute.err: %v", err)
	}
}

//服务端流
func printLists(client pd.StreamServiceClient, r *pd.StreamRequest) error {
	stream, err := client.List(context.TODO(), r)
	if err != nil {
		log.Panicln(err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Panicln(err)
		}

		log.Println(resp.Pt.Name, resp.Pt.Value)
	}

	return nil
}

//客户端流
func printRecord(client pd.StreamServiceClient, r *pd.StreamRequest) error {
	stream, err := client.Record(context.TODO())
	if err != nil {
		log.Panicln(err)
	}

	for n := 0; n <= 10; n++ {
		stream.Send(&pd.StreamRequest{
			Pt: &pd.StreamPoint{
				Name:  r.Pt.Name,
				Value: int64(n),
			},
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Panicln(err)
	}

	log.Println(res.Pt.Name, res.Pt.Value)
	return nil
}

//双向流
func printRoute(client pd.StreamServiceClient, r *pd.StreamRequest) error {
	stream, err := client.Route(context.TODO())
	if err != nil {
		log.Panicln(err)
	}

	go func() {
		for {
			stream.Send(&pd.StreamRequest{
				Pt: &pd.StreamPoint{
					Name:  r.Pt.Name,
					Value: int64(rand.Int()),
				},
			})

			time.Sleep(time.Millisecond * 100)
		}

	}()

	go func() {
		for {
			r, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(r.Pt.Name, r.Pt.Value)
		}
	}()

	time.Sleep(time.Second)

	err = stream.CloseSend()
	if err != nil {
		log.Panicln(err)
	}
	return nil
}
