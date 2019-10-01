package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"os"
)

func HandlecError(err error, when string)  {
	if err != nil{
		fmt.Println("error happen at", when)
		os.Exit(500)
	}
}

func main()  {
	// 连接redis操作
	conn, err := redis.Dial("tcp","127.0.0.1:6379")
	HandlecError(err, "connect")
	defer func() {
		_ = conn.Close()
	}()

	// set字符串
	_,_ = conn.Do("set","name", "wxz")

	// get字符串
	reply, err := conn.Do("get", "name")
	HandlecError(err, "doing command")

	fmt.Printf("type=%T， value=%s\n", reply, reply)

	nameStr,_ := redis.String(reply, err)
	//ageStr,_ := redis.Int(reply, err)
	 fmt.Println(nameStr)

	// hash
	_,_ = conn.Do("hmset", "person", "name", "hxh", "age", 21)
	hsh,err1 := conn.Do("hgetall", "person")
	if err1 != nil{
		fmt.Println("错误了：", err1)
	}
	person,_ := redis.String(hsh, err1)
	fmt.Println(person)






	
}
