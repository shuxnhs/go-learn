package main

import "fmt"

func add(a int)int  {
	a = a + 1
	return a
}
// 传指针
func addpoint(b *int)int  {
	*b = *b + 1
	return *b
}

func main()  {
	x := 3
	fmt.Print(" x = ",  x)
	x1 := add(x)
	fmt.Print(" x+1 = ", x1)
	// add()中传递的是x的copy，所以x本身不会加一
	fmt.Print(" x = ", x)

	y := 3
	fmt.Print(" y = ", y)
	y1 := addpoint(&y)
	fmt.Print(" y + 1 =", y1)
	// y本身加一了，addpoint传递的是指针，传指针可以使多个函数操作同一个对象
	fmt.Print(" y =", y)
}
