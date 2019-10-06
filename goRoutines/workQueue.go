package main

import (
	"fmt"
	"time"
)

func work(count int, name string, workQueue chan string)  {
	for i := 1; i <= count; i++ {
		fmt.Println(name, "输出：", i)
		time.Sleep(1 * time.Second)
	}
	workQueue <- name+ "finish"
}

/**
 * @desc： 管道的调度例子
 */
func main()  {
	// 如果没有管道调度，主协程执行完毕就会结束,不会管子协程
	workQueue := make(chan string, 3)
	go work(10, "子协程1", workQueue)

	work(5,"主协程", workQueue)

	// 通过管道的读取来阻塞主协程
	for i := 0; i < 2 ; i++ {
		<-workQueue
	}

	fmt.Println("mainTask finish")

}
