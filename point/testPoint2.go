package main

import "fmt"

func swap1(a, b int) {
	a, b = b, a
}

func swap2(a, b *int) {
	a, b = b, a
	fmt.Println("swap2交换过程中：", *a, *b)
}

// *取值 &取地址
func swap3(a, b *int) {
	t := *a
	*a = *b
	*b = t
	fmt.Println("swap2交换过程中：", *a, *b)
}

func main() {
	a, b := 1, 2
	swap1(a, b)
	fmt.Println(a, b)
	c, d := 1, 2
	swap2(&c, &d)
	fmt.Println(c, d)
	e, f := 1, 2
	swap3(&e, &f)
	fmt.Println(e, f)
}
