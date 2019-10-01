package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func DoCommand(pool *redis.Pool, i int)  {
	conn := pool.Get()
	defer conn.Close()
	// 开始执行redis操作
	replay, err := conn.Do("get", "name")
	name,_ := redis.String(replay,err)
	fmt.Println("name: ", name, "\n")
}

func main()  {

	// 新建一个redis连接池
	pool := &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
		TestOnBorrow:    nil,
		MaxIdle:         20,	// 最大闲置数
		MaxActive:       0,		// 最大连接数，0表示无限
		IdleTimeout:     100 * time.Second,		// 闲置连接的超时时间
		Wait:            false,
		MaxConnLifetime: 0,		// 允许最长连接时间，0表示无限
	}

	// 延时关闭连接
	defer pool.Close()

	// 开启并发连接
	for i:= 0; i < 10; i++{
		go DoCommand(pool, i)
	}

	// 保持主协程存活
	time.Sleep(3 * time.Second)
}
