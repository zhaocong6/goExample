package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	//pub()
	use()
}

func conn() *amqp.Connection {
	var uri = fmt.Sprintf("amqp://%s:%s@%s:%d", "test", "test", "172.16.40.185", 5672)

	//创建一个mq连接
	conn, err := amqp.Dial(uri)
	if err != nil {
		log.Panic(err)
		return nil
	}

	return conn
}

func channel(conn *amqp.Connection) *amqp.Channel {
	channel, err := conn.Channel()
	if err != nil {
		log.Panic(err)
		return nil
	}

	return channel
}

func pub() {
	channel := channel(conn())
	//声明一个test的直连交换机
	//持久化
	err := channel.ExchangeDeclare("testExchange", "direct", true, false, false, false, nil)
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
	channel := channel(conn())

	//声明一个队列
	q, err := channel.QueueDeclare("match", true, false, false, false, nil)
	if err != nil {
		log.Panic(err)
		return
	}

	//注册消费者
	msg, err := channel.Consume(q.Name, "match", true, false, false, false, nil)
	if err != nil {
		log.Panic(err)
		return
	}

	forever := make(chan bool)
	go func() {

		for {
			select {
			case res := <-msg:
				fmt.Println(string(res.Body))
			}
		}
	}()
	log.Printf("Waiting for messages. To exit press CTRL+C")
	<-forever
}
