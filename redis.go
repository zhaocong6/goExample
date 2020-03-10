package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var redisPool redis.Pool

func init() {
	redisPool = redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
		MaxIdle:      10,
		MaxActive:    100,
		IdleTimeout:  60,
		Wait:         true,
	}
}

func main() {
	pool()
	notPool()
}

func pool() {
	conn := redisPool.Get()
	defer conn.Close()

	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error:", err)
	}
	defer conn.Close()

	test, err := redis.String(conn.Do("GET", "test"))
	fmt.Println(test)
}

func notPool() {
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
