package main

import (
	"io/ioutil"
	"net/http"
)

func main()  {
	http.HandleFunc(
		"/getMyInfo",
		func(writer http.ResponseWriter, request *http.Request) {
			 _,_ = writer.Write([]byte("请求地址为"+request.Host+"\n"))
			 _,_ = writer.Write([]byte("请求远程地址为"+request.RemoteAddr+"\n"))
			 _,_ = writer.Write([]byte("请求协议为"+request.Proto+"\n"))
			 _,_ = writer.Write([]byte("请求方法为"+request.Method+"\n"))
			 _,_ = writer.Write([]byte("请求路由为"+request.RequestURI+"\n"))
			 _,_ = writer.Write([]byte("请求ua为"+request.UserAgent()+"\n"))
		})
	http.HandleFunc(
		"/getMyCSDN",
		func(writer http.ResponseWriter, request *http.Request) {
			csdn,err := ioutil.ReadFile("/usr/local/var/go/src/go-learn/network/http/csdn.html")
			if err == nil{
				_,_ = writer.Write(csdn)
			}else {
				_,_ = writer.Write([]byte("获取失败"))
			}

		})
	_ = http.ListenAndServe("127.0.0.1:8088", nil)
}