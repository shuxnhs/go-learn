package main

import (
	"flag"
	"fmt"
)

func main()  {
	// 返回指针
	age := flag.Int("age", 22, "你的年龄")
	name := flag.String("name", "hxh", "你的名字")
	// 解析参数
	flag.Parse()
	// 返回的是指针
	//	func String(name string, value string, usage string) *string {
	//		return CommandLine.String(name, value, usage)
	//	}
	fmt.Println(*age, *name)
}