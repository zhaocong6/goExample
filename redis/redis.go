package main

import (
	"fmt"
	red "github.com/gomodule/redigo/redis"
	"time"
)

type Redis struct {
	pool *red.Pool
}

var redis = &Redis{}

func init() {
	redis.pool = &red.Pool{
		MaxIdle:         256,  //最大空闲连接数
		MaxActive:       5000, //最大连接数
		IdleTimeout:     time.Duration(120),
		Wait:            true,
		MaxConnLifetime: 0,
		Dial: func() (conn red.Conn, err error) {
			return red.Dial(
				"tcp",
				"127.0.0.1:6379",
				red.DialReadTimeout(time.Duration(1000)*time.Millisecond),
				red.DialWriteTimeout(time.Duration(1000)*time.Millisecond),
				red.DialConnectTimeout(time.Duration(1000)*time.Millisecond),
				red.DialDatabase(0),
				red.DialPassword(""),
			)
		},
	}
}

func main() {
	con := redis.pool.Get()
	fmt.Println(con.Do("set", "test", "测试"))
	fmt.Println(red.String(con.Do("get", "test")))
}
