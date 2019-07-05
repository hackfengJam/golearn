package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		>>>
			打印后
			打印中
			打印前
			panic: 触发异常
				panic: 触发异常后
				panic: 触发异常中
				panic: 触发异常前
	*/
	defer_call()
}

func defer_call() {
	defer func() {
		fmt.Println("打印前")
		time.Sleep(1 * time.Second)
		panic("触发异常前")
	}()
	defer func() {
		fmt.Println("打印中")
		time.Sleep(1 * time.Second)
		panic("触发异常中")

	}()
	defer func() {
		fmt.Println("打印后")
		time.Sleep(1 * time.Second)
		panic("触发异常后")
	}()

	panic("触发异常")
}
