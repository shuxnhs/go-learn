package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

// 全局变量申明
var (
	store    = "hxh and wxz"
	together = 100
)

func main() {
	fmt.Println("vim-go")

	// 1.声明类型,int默认值为0
	var one int
	fmt.Println("one = ", one)

	// 2.由go自行判断变量类型
	var two = 520
	fmt.Println("two =", two)

	// 3.直接赋值
	three := 13.14
	fmt.Println("three =", three)

	// 4.多个变量赋值
	var boy, age = "hxh", 10
	fmt.Println("boy = ", boy, "age = ", age)

	// 5. 多个变量使用类型推到
	girl, girlage := "wxz", 11
	fmt.Println("girl =", girl, "age = ", girlage)

	fmt.Println("store is", store, "together with", together)

	// 6. 查看变量类型以及占用的空间数
	fmt.Printf("boy age 的类型是: %T", age, "占用空间大小为: %d", unsafe.Sizeof(age))

	// 7.字符本质是一个整数,就是字符对应的码值
	var testvar int
	testvar = '何'
	fmt.Printf("\n testvar本质是%d%c", testvar, testvar)

	// 7.字符串,字符串一旦赋值则无法再改变, 只是没办法修改story[0], story[2]其中的字符
	var story string = "wxz❤️ hxh"
	fmt.Println("\n", story)
	fmt.Printf("\nstory is %T", story)

	// 8.字符串使用反引号可以原样输出
	var testenter string = `enter\n\r\tenter`
	fmt.Println("原样输出", testenter)

	testenter = `testnew`
	fmt.Println("change: ", testenter)

	// 强制类型转换为string
	testchang := fmt.Sprintf("%d", age)
	fmt.Printf("\nage type change %T%d", testchang, testchang)

	// 第二种方式,使用strconv包
	testchange2 := strconv.FormatInt(int64(girlage), 10)
	fmt.Println("testchange2", testchange2)

	//  (使用strconv)
	changeint := strconv.Itoa(age)
	fmt.Println(changeint)

}
