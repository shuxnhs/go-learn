package main

import (
	"errors"
	"fmt"
	"math"
)

// panic处理
func calVolumn(r float64) float64  {
	if r < 0{
		panic("传了个负数")
	}
	return 4/3*math.Pi*math.Pow(r, 3)
}

// 没想到吧，我还可以给你个error
func calVolumnVersion(r float64) (v float64, err error)  {
	if r < 0{
		err = errors.New("没想到吧，你给我传个负数，我给你抛个错误")
		return
	}
	return 4/3*math.Pi*math.Pow(r, 3),nil
}

// 自定义处理
type InvalidError struct {
	minRadius int
}

func (ine *InvalidError)dealError() string{
	// sprint不输出来只用来赋值
	info := fmt.Sprint("最小为", ine.minRadius)
	return info
}

func calVolumnVersionHigh(r float64) (v float64, err error)  {
	if r < 0{
		err = errors.New("没想到吧，你给我传个负数，我给你抛个错误")
		return
	}
	return 4/3*math.Pi*math.Pow(r, 3),nil
}

func main()  {

	v2,err1 := calVolumnVersion(-3)
	if err1 == nil{
		fmt.Println(v2)
	}else {
		fmt.Println("有问题：", err1)
	}

	// 让你panic前复活
	defer func() {
		if err:= recover(); err != nil{
			fmt.Println("找到原因了： \"", err, "\"")
			fmt.Println("你居然想传负数来害我,还好我recover了")
		}
	}()
	v := calVolumn(-3)
	fmt.Println(v)
}
