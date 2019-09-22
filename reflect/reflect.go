package main

import (
	"fmt"
	"reflect"
)

// 反射

type People struct {
	name string
	age int
	hobby []string
}

func (p *People)GetName()  {
	fmt.Println("名字是", p.name, "开始吃饭")
}

func (p *People)GetAge()  {
	fmt.Println("今年", p.age, "岁")
}

type Men struct {
	// 继承类people
	People
	hair bool
}

func (man *Men)GetHair()  {
	if man.hair != true{
		fmt.Println("男人的头发是短的")
	}else {
		fmt.Println("男人的头发是长的")
	}
}


func main()  {
	//people := People{"hxh", 11, []string{"吃","喝", "拉", "撒"}}
	// 新的创建方法
	man := new(Men)
	man.People = People{name:"wxz", age: 11}
	valueAPI(man)
}

func valueAPI(o Men)  {
	oValue := reflect.ValueOf(o)

	fmt.Println(oValue)
	fmt.Println("查看原始类型", oValue.Kind())

	fmt.Println("打印出所有属性的值")
	for i := 0; i < oValue.NumField(); i++ {
		fValue := oValue.Field(i)
		fmt.Println(fValue.Interface())		// 获得值的正射
	}

	fmt.Println("获取父类属性的值")
	fvalue := oValue.FieldByIndex([]int{0,0})
	fmt.Println(fvalue.Interface())
}
