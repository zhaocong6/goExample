package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error:", err)
	}
	defer conn.Close()

	test, err := redis.String(conn.Do("GET", "test"))
	fmt.Println(test)

	_, err = conn.Do("set", "test", "test")

	test, err = redis.String(conn.Do("GET", "test"))
	fmt.Println(test)
}
