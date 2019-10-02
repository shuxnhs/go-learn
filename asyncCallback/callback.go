package main

import (
	"fmt"
	"time"
)

func AfterWork()  {
	fmt.Println("WORK FINISH!")
}

func main()  {
	go func(f func()) {
		for i:=0; i < 9; i++ {
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
		f()
	}(AfterWork)

	time.Sleep(10 * time.Second)
}
