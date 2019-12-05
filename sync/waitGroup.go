package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * @desc 等待组
 */
func main() {
	// 等待组的三个API： Add wait done
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("子协程1开始了")
		time.Sleep(time.Second * 3)
		// 子协程结束了等待组就调用done
		wg.Done()
		fmt.Println("子协程1结束了")
	}()

	wg.Add(1)
	go func() {
		fmt.Println("子协程2开始了")
		// 使用timer
		<-time.After(4 * time.Second)
		//timer := time.NewTimer(4 * time.Second)
		//<- timer.C
		wg.Done()
		fmt.Println("子协程2结束了")
	}()

	wg.Add(1)
	go func() {
		fmt.Println("子协程3开始了")
		// 使用kicker
		ticker := time.NewTicker(1 * time.Second)
		for i := 0; i < 5; i++ {
			<-ticker.C
		}
		wg.Done()
		fmt.Println("子协程3结束了")
	}()

	// 三个子协程结束main函数就结束，使用等待组可以准确结束不用写for循环等待
	wg.Wait()
	fmt.Println("main over")
}
