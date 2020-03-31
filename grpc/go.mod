module goExample/grpc

go 1.13

replace github.com/zhaocong6/proto/hello => ./proto/hello

require (
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/zhaocong6/proto/hello v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.28.0
)
