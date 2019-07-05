package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func channel_test() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}

		// 在这里才合理
		//close(ch)
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	// panic: send on closed channel
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}
func main() {
	SetValue(21)
}

var value int32

func SetValue(delta int32) {
	/*
		TODO
		在前面示例的for循环中，我们使用语句v := value为变量v赋值。
		但是，要注意，其中的读取value的值的操作并不是并发安全的。
		在该读取操作被进行的过程中，其它的对此值的读写操作是可以被同时进行的。它们并不会受到任何限制。
		在第7章的第1节的最后，我们举过这样一个例子：在32位计算架构的计算机上写入一个64位的整数。
		如果在这个写操作未完成的时候有一个读操作被并发的进行了，那么这个读操作很可能会读取到一个只被修改了一半的数据。
		这种结果是相当糟糕的。
		为了原子的读取某个值，sync/atomic代码包同样为我们提供了一系列的函数。这些函数的名称都以“Load”为前缀，意为载入。
		我们依然以针对int32类型值的那个函数为例。
	*/
	for {
		v := value
		/*
			http://ifeve.com/go-concurrency-atomic/
			v := atomic.LoadInt32(&value)
		*/
		if atomic.CompareAndSwapInt32(&value, v, v+delta) {
			break
		}
	}
}
