package main

import (
	"context"
	"fmt"
	"time"
)

func work(ctx context.Context, msg string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value(msg))
			println(msg, "goroutinue is finish......")
			return
		default:
			println("goroutinue is running", time.Now().String())
			time.Sleep(time.Second)
		}
	}

}

func main() {

	// simple
	ctx, cancel := context.WithCancel(context.Background())
	go work(ctx, "withCancel")
	time.Sleep(time.Second * 3)
	println("cancel......")
	cancel()
	time.Sleep(time.Second * 3)
	println("finish")

	// withvalue
	ctx1, valueCancel := context.WithCancel(context.Background())
	valueCtx := context.WithValue(ctx1, "key1", "value1")
	go work(valueCtx, "key1")
	time.Sleep(time.Second * 3)
	valueCancel()
	time.Sleep(time.Second * 3)
	println("finish")
}
