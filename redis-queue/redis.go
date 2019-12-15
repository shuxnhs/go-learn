package main

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
)

type Producer struct {
	// 生产者
}

func (p *Producer) publish(conn redis.Conn, listKey string, data string) (reply interface{}, err error) {
	return conn.Do("lpush", listKey, data)
}

type Customer struct {
	// 消费者
}

func (c *Customer) putMessage(conn redis.Conn, listKey string) (interface{}, error) {
	return conn.Do("rpop", listKey)
}

func (c *Customer) getCount(conn redis.Conn, listKey string) (interface{}, error) {
	return conn.Do("llen", listKey)
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
	// 连接redis操作
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	HandlecError(err, "connect")
	defer func() {
		_ = conn.Close()
	}()

	producer := Producer{}
	personMap := make(map[string]interface{})
	personMap["name"] = "hxh"
	personMap["work"] = "toDoSomething"
	bytes, _ := json.Marshal(personMap)
	_, _ = producer.publish(conn, "test_queue", string(bytes))

	customer := Customer{}
	num, _ := customer.getCount(conn, "test_queue")
	fmt.Println("队列数量为", num)

	values, err := redis.String(customer.putMessage(conn, "test_queue"))
	dataMap := make(map[string]interface{}) // 准备好map来装
	_ = json.Unmarshal([]byte(values), &dataMap)
	fmt.Println(dataMap["work"])

}
