package main

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"os"
	"sync"
	"time"
)

type delayQueue struct {
	// 延迟队列
}

func (d *delayQueue) publish(conn redis.Conn, zSetKey string, dataMap map[string]interface{}, time int64) (reply interface{}, err error) {
	// 生成唯一id，保证zset的每一个value都不一样,time为执行的时间戳
	dataMap["uuid"] = uuid.NewV4().String()
	bytes, _ := json.Marshal(dataMap)
	return conn.Do("zadd", zSetKey, time, string(bytes))
}

func (d *delayQueue) customer(conn redis.Conn, zSetKey string) {
	for true {
		now := time.Now().Unix()
		data, err := redis.Strings(conn.Do("zrangebyscore", zSetKey, 0, now, "limit", 0, 1))
		if err == nil && len(data) > 0 {
			res, delErr := conn.Do("zrem", "test-delay-queue", data[0])
			if res.(int64) >= 1 && delErr == nil {
				dataMap := make(map[string]interface{}) // 准备好map来装
				_ = json.Unmarshal([]byte(data[0]), &dataMap)
				fmt.Println("任务是：", dataMap["work"])
			} else {
				fmt.Println(delErr)
			}
		} else {
			time.Sleep(time.Second * 10)
			continue
		}
	}
}

func HandlecError(err error, when string) {
	if err != nil {
		fmt.Println("error happen at", when)
		os.Exit(500)
	} else {
		fmt.Println("连接成功")
	}
}

func main() {
	var wg sync.WaitGroup

	// 连接redis操作
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	HandlecError(err, "connect")
	defer func() {
		_ = conn.Close()
	}()

	delayQueue := delayQueue{}
	personMap := make(map[string]interface{})
	personMap["name"] = "hxh"
	personMap["work"] = "toDoSomething"
	_, _ = delayQueue.publish(conn, "test-delay-queue", personMap, time.Now().Unix())

	wg.Add(1)
	go func() {
		delayQueue.customer(conn, "test-delay-queue")
		wg.Done()
	}()
	wg.Wait()

}
