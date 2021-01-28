package main

import (
	"fmt"
	"runtime"
)

func main() {
	/*
		https://yushuangqi.com/blog/2017/golang-mian-shi-ti-da-an-yujie-xi.html?from=groupmessage&isappinstalled=0

		有可能触发异常，是随机事件。

		随机，可能报错，也可能不报错
	*/

	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}
