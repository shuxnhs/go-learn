package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * @desc：读写锁/互斥锁
 * 多路可读，一路可写
 */

func main() {
	var wg sync.WaitGroup
	var wr sync.RWMutex

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			// 加读锁
			wr.RLock()
			fmt.Println("假装开始读数据库, select * from table")
			<-time.After(3 * time.Second)
			wr.RUnlock()
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			// 写数据库，每次只能有一个写,加写锁
			wr.Lock()
			fmt.Println("假装写入数据库， insert into table")
			<-time.After(3 * time.Second)
			wr.Unlock()
			wg.Done()
		}()

	}
	wg.Wait()
}
