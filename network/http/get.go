package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func HandlecClientError(err error, when string) {
	if err != nil {
		fmt.Println("error happen at", when)
		os.Exit(500)
	}
}

/**
 * @desc : 发送get请求
 */
func GetRequest(url string) string {
	resp, err := http.Get(url)
	HandlecClientError(err, "request get")
	defer func() {
		_ = resp.Body.Close()
	}()
	// ioutil去读body(io.ReadCloser)
	bytes, _ := ioutil.ReadAll(resp.Body)
	return string(bytes)
}

func GoTest(url string) {
	go func() {
		for i := 0; i < 1000; i++ {
			response := GetRequest(url)
			fmt.Println(response)
		}
	}()
}

func main() {
	// 读取器，获取终端输入的URL去请求网页
	reader := bufio.NewReader(os.Stdin)
	lineByte, _, err := reader.ReadLine()
	HandlecClientError(err, "readLine")
	url := string(lineByte)
	num := 100
	go func(num int, url string) {
		for i := 0; i < num; i++ {
			GoTest(url)
		}
	}(num, url)

	time.Sleep(1000000 * time.Second)
}
