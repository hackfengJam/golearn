// Package pubsub implements a simple multi-topic pub-sub library.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 生产者：生成 factor 整数倍的序列
func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

// 消费者
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func SimplePS() {
	ch := make(chan int, 64) // 成果队列

	go Producer(3, ch) // 生成 3 的倍数的序列
	go Producer(5, ch) // 生成 5 的倍数的序列
	go Consumer(ch)    // 消费 生成的队列

	// 运行一定时间后退出
	time.Sleep(5 * time.Second)
	/*
		我们开启了2个Producer生产流水线，分别用于生成3和5的倍数的序列。然后开启1个Consumer消费者线程，打印获取的结果。
		我们通过在main函数休眠一定的时间来让生产者和消费者工作一定时间。正如前面一节说的，这种靠休眠方式是无法保证稳定的输出结果的。
	*/
}

func main() {
	ch := make(chan int, 64) // 成果队列

	go Producer(3, ch) // 生成 3 的倍数的序列
	go Producer(5, ch) // 生成 5 的倍数的序列
	go Consumer(ch)    // 消费 生成的队列

	// Ctrl+C 退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
	/*
		我们这个例子中有2个生产者，并且2个生产者之间并无同步事件可参考，它们是并发的。
		因此，消费者输出的结果序列的顺序是不确定的，这并没有问题，生产者和消费者依然可以相互配合工作。
	*/
}
