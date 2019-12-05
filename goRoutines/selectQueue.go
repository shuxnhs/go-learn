package main

import (
	"fmt"
	"time"
)

/**
 * @desc: select随机调用能执行的
 */
func main() {
	chA := make(chan int, 5)
	chB := make(chan int, 4)
	chC := make(chan int, 3)

	chB <- 1
	chB <- 2
	chB <- 3
	chB <- 4
	chC <- 1
	chC <- 2
	chC <- 3

OUTER:
	for {
		select {
		case chA <- 123:
			fmt.Println("执行任务A")
			time.Sleep(time.Second)
		case b := <-chB:
			fmt.Println("执行任务B", b)
			time.Sleep(time.Second)
		case c := <-chC:
			fmt.Println("执行任务C", c)
			time.Sleep(time.Second)
		default:
			fmt.Println("任务结束")
			break OUTER
		}
	}
}
