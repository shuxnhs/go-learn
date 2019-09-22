package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func HandlecClientPostError(err error, when string)  {
	if err != nil{
		fmt.Println("error happen at", when)
		os.Exit(500)
	}
}

/**
 * @desc : 发送post请求
 */
func PostRequest(url string, contentType string, reader io.Reader) string {
	resp, err := http.Post(url, contentType, reader)
	HandlecClientPostError(err, "request post")
	defer func() {
		_ = resp.Body.Close()
	}()
	// ioutil去读body(io.ReadCloser)
	bytes, _ := ioutil.ReadAll(resp.Body)
	return string(bytes)
}

func main()  {
	// 读取器，获取终端输入的URL去请求网页
	reader := bufio.NewReader(os.Stdin)
	//url,_,err := reader.ReadLine()
	//HandlecClientPostError(err, "read url")
	fmt.Println("please input the params")
	url := "https://httpbin.org/post"
	params,_,err := reader.ReadLine()	// name=hxh&age=21
	HandlecClientPostError(err, "read params")
	contentType := "application/x-www-form-urlencoded"
	response := PostRequest(string(url), contentType, strings.NewReader(string(params)))
	fmt.Println(response)
}