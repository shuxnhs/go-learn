package main

import (
	"fmt"
	"net"
	"os"
)

func HandleError(err error, when string)  {
	if err != nil{
		fmt.Println("error happen at", when)
		os.Exit(500)
	}
}

func ioHandleConnect(conn net.Conn)  {
	// 准备缓冲区
	buffer := make([]byte, 1024)

	// 开始读取信息
	for   {
		n, err := conn.Read(buffer)		// n
		HandleError(err, "read buffer")
		clientMsg := string(buffer[:n])
		fmt.Println("received from ",conn.RemoteAddr(),": ",clientMsg)
		if clientMsg == "exit"{
			_,_ = conn.Write([]byte("bye~~"))
			break
		}
		_,_ = conn.Write([]byte("receive"))
	}

	fmt.Println("over~")
}

func main()  {
	listener, err := net.Listen("tcp","127.0.0.1:8811")
	HandleError(err,"LISTEN")
	fmt.Println("listening")

	for {
		conn, e := listener.Accept()
		HandleError(e, "connect")

		// 开启一条协程去处理连接
		go ioHandleConnect(conn)
	}
}
