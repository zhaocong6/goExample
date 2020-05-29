module goExample

go 1.13

require (
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/samuel/go-zookeeper v0.0.0-20190923202752-2cc03de413da // indirect
	github.com/spf13/cobra v1.0.0 // indirect
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
	github.com/zhaocong6/goUtils v1.0.3
	go.mongodb.org/mongo-driver v1.3.3 // indirect
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/zhaocong6/proto/hello => ./Grpc/proto/hello
