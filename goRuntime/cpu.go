package main

import (
	"fmt"
	"runtime"
)

func main()  {
	fmt.Println("当前可用的cpu逻辑核心数有：", runtime.NumCPU())
	fmt.Println("修改为4前的逻辑cpu逻辑核心数有：", runtime.GOMAXPROCS(4))
	fmt.Println("修改为2前的逻辑cpu逻辑核心数有：", runtime.GOMAXPROCS(2))
}
