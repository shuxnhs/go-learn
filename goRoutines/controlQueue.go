package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

/**
 * @desc 用管道控制并发数
 */

func countSqrt(name string, n int, sem chan string) {
	// 每次调用前先写入管道，没有写入管道的就无法调用
	sem <- name
	ret := math.Sqrt(float64(n))
	fmt.Println(n, "的平方根为：", ret)
	time.Sleep(time.Second)
	// 调用结束从管道释放
	<-sem
}

func main() {
	// 用管道（信号量）去控制并发数
	semphore := make(chan string, 5)
	for i := 0; i < 100; i++ {
		go countSqrt("协程"+strconv.Itoa(i), i, semphore)
	}
	for {
		time.Sleep(time.Second)
	}
}
