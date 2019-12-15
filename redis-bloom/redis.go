package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
)

func HandlecError(err error, when string) {
	if err != nil {
		fmt.Println("error happen at", when)
		os.Exit(500)
	} else {
		fmt.Println("连接成功")
	}
}

func main() {
	// 连接redis操作
	conn, err := redis.Dial("tcp", "127.0.0.1:6380")
	HandlecError(err, "connect")
	defer func() {
		_ = conn.Close()
	}()

	// 测试redis的布隆过滤器
	res, err := conn.Do("bf.add", "codehole", "user1")
	fmt.Println(res, err)

	res2, err2 := conn.Do("bf.exists", "codehole", "user1")
	fmt.Println(res2, err2)
}
