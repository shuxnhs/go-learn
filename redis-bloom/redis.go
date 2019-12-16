package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"os"
	"strconv"
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
	res, err := conn.Do("bf.add", "codehole", "user100")
	fmt.Println(res, err)

	res2, err2 := conn.Do("bf.exists", "codehole", "user1")
	fmt.Println(res2, err2)

	testKey := uuid.NewV4().String()

	// 对已经添加进布隆过滤器的他肯定是能判断出来是存在的（res=1）
	for i := 1; i < 100000; i++ {
		_, _ = conn.Do("bf.add", testKey, "user"+strconv.Itoa(i))
		res, _ = conn.Do("bf.exists", testKey, "user"+strconv.Itoa(i))
		if res.(int64) == 0 {
			fmt.Println(i)
			break
		}
	}

	// 对于没添加进布隆过滤器的，他可能会误判为存在（res=1）
	for i := 1; i < 100000; i++ {
		_, _ = conn.Do("bf.add", testKey, "user"+strconv.Itoa(i))
		res, _ = conn.Do("bf.exists", testKey, "user"+strconv.Itoa(i+1))
		if res.(int64) == 1 {
			fmt.Println(i)
			break
		}
	}
}
