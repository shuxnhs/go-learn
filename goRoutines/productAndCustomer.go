package main

import (
	"fmt"
	"strconv"
	"time"
)

/**
 * @desc 管道的生产者消费者模型
 */

// 1.商品的类
type product struct{
	name string
}

/**
 * @desc: 生产函数
 * @channel productList 商品只写管道，只允许写入
 */
func produce(productList chan <- product)  {
	for  {
		product := product{"产品" + strconv.Itoa(time.Now().Second())}
		productList <- product
		fmt.Println("生产了" , product)
		time.Sleep(1 * time.Second)
	}
}

/**
 * @desc：消费函数
 * @channel productList 商品只读管道，只允许读入
 * @channel countList   次数只读管道，只允许写入
 */
func customer(productList <- chan product, countList chan <- int)  {
	for  {
		product := <- productList
		fmt.Println("消费掉：", product)
		countList <- 1
	}
}

func main()  {
	// 新建商品管道与计数管道
	productList := make(chan product, 5)
	countList := make(chan int, 10)

	go produce(productList)

	go customer(productList, countList)

	for i := 0; i < 10; i++ {
		<- countList
	}

	fmt.Println("finish")

}
