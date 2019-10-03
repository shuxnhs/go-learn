package singleton

import (
	"fmt"
	"sync"
)

// 全局实例
type singleton struct {
	data int
}

// 小写私有实例变量
var sing *singleton

// 保证线程安全，只执行一次
var once sync.Once

func GetInstance(num int) *singleton {
	once.Do(func() {
		sing = &singleton{data:num}
		fmt.Println("实例对象的值为和地址为:", sing, &sing)
	})
	return sing
}
