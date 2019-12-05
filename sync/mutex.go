package main

import (
	"fmt"
	"sync"
)

/**
 * 同步锁/资源锁
 */

func main() {

	var mt sync.Mutex
	var Amoney = 1000
	var Bmoney = 1000
	var mg sync.WaitGroup

	// 没有加同步锁的情况
	for i := 0; i < 10; i++ {
		mg.Add(1)
		go func() {
			for j := 0; j < 100000; j++ {
				Amoney += 1
			}
			mg.Done()
		}()
	}

	// 加上锁，保证并发安全
	for a := 0; a < 10; a++ {
		mg.Add(1)
		go func(index int) {
			fmt.Println("协程", index, "开始抢锁")
			mt.Lock()
			fmt.Println("为协程", index, "加上锁")
			for b := 0; b < 100000; b++ {
				Bmoney += 1
			}
			mt.Unlock()
			fmt.Println("执行完毕，协程", index, "解锁")
			mg.Done()
		}(a)
	}
	mg.Wait()
	fmt.Println("没有加同步锁的钱数：", Amoney)
	fmt.Println("加同步锁后的钱数", Bmoney)
	fmt.Println("正确结果为：1000 + 10 * 100000 = ", 1000+10*100000)
}
