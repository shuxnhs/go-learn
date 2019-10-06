package main

import (
	"fmt"
	"time"
)

func main()  {
	var channel chan int
	fmt.Println(channel)

	// make只是创建了一个底层数据结构的引用，当赋值或参数传递的时候只是拷贝了一个引用
	channel = make(chan int)
	fmt.Println(channel)

	// 读写基本操作，必须先初始化后才能进行，否则会死锁；管道已满，还继续写入或读出会发生死锁
	channel = make(chan int, 2)
	channel <- 1
	fmt.Println(<- channel)

	// 持续的写入，必须有协程持续读出
	fmt.Println("开始持续的写入......")
	go func() {
		// 开始读出
		for i := 0; i < 10 ; i++ {
			fmt.Println(<-channel)
		}
	}()
	for j := 0; j < 10; j++ {
		channel <- j
		time.Sleep(1 * time.Second)
	}

	// 关闭管道,无法再继续进行管道的写入，但可以继续读出数据
	// 关闭一个未初始化的管道会panic
	close(channel)
	num, ok := <-channel
	if ok {
		fmt.Println(num)
	}else {
		fmt.Println("管理里没有数据了",num ,ok)
	}

	// 关闭管道会通知所有的协程，取消继续的阻塞运行
	myChan := make(chan string, 3)


	myChan <- "hxh"
	fmt.Println("管道的长度：", len(myChan))
	fmt.Println("管道的容量：", cap(myChan))
	// 已写满的管道继续写入或已为空的管道继续读出会使主协程阻塞，死锁
	myChan <- "wxz"
	time.Sleep(2 * time.Second)
	go func() {
		for x := range myChan {
			fmt.Println(x)
		}
		fmt.Println("子协程结束")
	}()
	// 关闭后，不回再阻塞读出
	close(myChan)


	
}
