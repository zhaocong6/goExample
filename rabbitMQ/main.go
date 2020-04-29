package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	pub()
	use()
}

func pub() {
	var uri = fmt.Sprintf("amqp://%s:%s@%s:%d", "guest", "guest", "127.0.0.1", 5672)

	//创建一个mq连接
	conn, err := amqp.Dial(uri)
	if err != nil {
		log.Panic(err)
		return
	}
	defer conn.Close()

	//创建一个channel
	channel, err := conn.Channel()
	if err != nil {
		log.Panic(err)
		return
	}
	defer channel.Close()

	//声明一个test的直连交换机
	//持久化
	err = channel.ExchangeDeclare("testExchange", "direct", true, false, false, false, nil)
	if err != nil {
		log.Panic(err)
		return
	}

	//声明一个队列
	_, err = channel.QueueDeclare("testQueue", true, false, false, false, nil)
	if err != nil {
		log.Panic(err)
		return
	}

	//给队列绑定route和交换机
	channel.QueueBind("testQueue", "testRoute", "testExchange", false, nil)

	err = channel.Publish("testExchange", "testRoute", false, false,
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte("hello world"),
		})

	if err != nil {
		log.Panic(err)
		return
	}

	log.Println("success")
}

func use() {
	var uri = fmt.Sprintf("amqp://%s:%s@%s:%d", "guest", "guest", "127.0.0.1", 5672)

	//创建一个mq连接
	conn, err := amqp.Dial(uri)
	if err != nil {
		log.Panic(err)
		return
	}
	defer conn.Close()

	//创建一个channel
	channel, err := conn.Channel()
	if err != nil {
		log.Panic(err)
		return
	}
	defer channel.Close()

	//声明一个队列
	q, err := channel.QueueDeclare("testQueue", true, false, false, false, nil)
	if err != nil {
		log.Panic(err)
		return
	}

	//注册消费者
	msg, err := channel.Consume(q.Name, "testConsume", true, false, false, false, nil)
	if err != nil {
		log.Panic(err)
		return
	}

	forever := make(chan bool)
	go func() {

		for {
			select {
			case res := <-msg:
				log.Println(res.Type)
				log.Println(res.MessageId)
				log.Printf("Received a message: %s", res.Body)
			}
		}
	}()
	log.Printf("Waiting for messages. To exit press CTRL+C")
	<-forever
}
