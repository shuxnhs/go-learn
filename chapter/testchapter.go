package main

import "fmt"

func main() {

	// go的除法
	num1 := 10
	fmt.Println("计算10/4 = ", num1/4)

	// 如果需要保留小数点,需要有浮点数来参与运算
	num2 := 10.0
	fmt.Println("计算10.0/4 = ", num2/4)

	// 求余公式  a % b = a - a / b * b
	fmt.Println(10 % 3)
	fmt.Println(-10 % 3)
	fmt.Println(10 % -3)
	fmt.Println(-10 % -3)

	// 自增运算符  ++ --
	num3 := 5
	fmt.Println(num3)
	// fmt.Println("num3++ = ", num3++)
	// fmt.Println("++num3 = ", ++num3)
	// ++ 只能作为一个独立语句,不能num++ 或 num2 := num++
	// ++只能放在语句后面
	num3++
	fmt.Println(num3)
}
