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

func ioHandleConnect(conn *net.UDPConn)  {
	// 准备缓冲区
	buffer := make([]byte, 1024)
	for  {
		n, remoteAddr,e := conn.ReadFromUDP(buffer)
		HandleError(e,"ReadFromUDP")

		clientMsg := string(buffer[:n])
		fmt.Println("reveived from :", remoteAddr,clientMsg)
		if clientMsg != "exit"{
			_,_ = conn.WriteToUDP([]byte("has reveived："+clientMsg),remoteAddr)
		}else {
			_,_ = conn.WriteToUDP([]byte("bye~"),remoteAddr)
		}

	}

}

func main()  {
	udp, err := net.ResolveUDPAddr("udp","127.0.0.1:8812")
	HandleError(err,"ResolveUDPAddr")
	fmt.Println("listening")

	// 开始去listen
	conn, e := net.ListenUDP("udp", udp)
	HandleError(e, "ListenUDP")

	// 直接去处理连接
	ioHandleConnect(conn)
}
