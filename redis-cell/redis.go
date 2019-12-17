package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"os"
	"time"
)

func HandlecError(err error, when string) {
	if err != nil {
		fmt.Println("error happen at", when)
		os.Exit(500)
	}
}
func main() {
	// 连接redis操作
	conn, err := redis.Dial("tcp", "127.0.0.1:6381")
	HandlecError(err, "connect")
	defer func() {
		_ = conn.Close()
	}()

	for i := 1; i < 10; i++ {
		// 漏斗限流测试
		key := "redis-cell"
		res, _ := conn.Do("CL.THROTTLE", key, 15, 30, 60)
		fmt.Println(res)
		time.Sleep(time.Millisecond * 500)
	}
}
