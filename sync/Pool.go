package main

import (
	"fmt"
	"runtime"
	"sync"
)

/**
 * 临时对象池
 */

func main() {

	pool := sync.Pool{New: func() interface{} {
		return 0
	}}
	pool.Put(1)
	a := pool.Get()
	fmt.Println(a)
	pool.Put(1)
	runtime.GC()
	b := pool.Get()
	fmt.Println(b)

}
