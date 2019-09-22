package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)


func HandlecClientError(err error, when string)  {
	if err != nil{
		fmt.Println("error happen at", when)
		os.Exit(500)
	}
}

func main()  {
	// 呼叫
	conn, e := net.Dial("udp", "127.0.0.1:8812")
	HandlecClientError(e,"dial")

	// 读取器,标准输入的读取器
	reader := bufio.NewReader(os.Stdin)
	buffer := make([]byte, 1024)

	// 阻塞读取输入然后发送给服务器
	for  {
		lineByte,_,err := reader.ReadLine()
		HandlecClientError(err, "readLine")

		// 开始写入
		_,_ = conn.Write(lineByte)

		// 开始获取服务端返回的信息
		n,_ := conn.Read(buffer)
		serverMsg := string(buffer[:n])
		fmt.Println("server say: ", serverMsg)

		if serverMsg == "bye~"{
			break
		}
		fmt.Println("over~~")
	}
}
