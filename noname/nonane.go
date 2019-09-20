package main

import (
	"fmt"
	"time"
)

func main(){

	// 4. 闭包函数
	Liubei := chuZhan("liubei", 1)
	Zhuge := chuZhan("zhuge",2)
	go func() {
		for i:= 0; i < 8 ; i++ {
			fmt.Println(Liubei())
			//time.Sleep(1 *time.Second)
		}
	}()

	go func() {
		for i:= 0; i < 8 ; i++ {
			fmt.Println(Zhuge())
			//time.Sleep(1 *time.Second)
		}
	}()


	// 1. 匿名函数与defer应用
	fmt.Println("输出第一列")
	defer func(){
		fmt.Println("输出倒数第二列")
		fmt.Println("输出倒数第一列")
	}()

	// 2. 并发
	// (1)并行运行在主协程
	fmt.Println("1")
	time.Sleep(1 * time.Second)
	fmt.Println("2")
	time.Sleep(1 * time.Second)

	// (2)开始并发。。运行在子协程，不影响下面的输出
	go func() {
		for i := 0; i< 10; i++ {
			// 下面的没办法输出完成，因为主协程已经结束。。除非主协程最后的sleep时间再长一点
			fmt.Println(i)
			time.Sleep(3 * time.Second)
		}
	}()

	fmt.Println("3")
	time.Sleep(1 * time.Second)
	fmt.Println("4")
	//time.Sleep(30 * time.Second)

	// 3. 有参数的匿名函数
	defer func(name string, age int)(grade int) {
		if name == "hxh" && age == 22 {
			grade = 100
		}else {
			grade = 0
		}
		fmt.Println(grade)
		return
	}("hxh", 22)

}


// 返回的是函数
func chuZhan(Leader string, startPos int) func() string  {
	var people = []string{"关羽", "张飞", "赵云", "马超", "黄忠"}
	var pos int = startPos
	chuZhanPeople := func() string{
		pos ++
		if pos > 4{
			pos = 0
		}
		return Leader + "派出" + people[pos]
	}
	return chuZhanPeople
}
