package main

import (
	"fmt"
	"sync"
)

func mutexFunc() {
	// 我们不能直接对一个未加锁状态的sync.Mutex进行解锁，这会导致运行时异常。
	// 下面这种方式并不能保证正常工作：
	//var mu sync.Mutex
	//
	//go func(){
	//	fmt.Println("你好, 世界")
	//	mu.Lock()
	//}()
	//
	//mu.Unlock()

	// 修复后的代码：
	var mu sync.Mutex

	mu.Lock()
	go func() {
		fmt.Println("Hello, world !")
		mu.Unlock()
	}()

	mu.Lock()
}

func chanFunc() {
	// 使用sync.Mutex互斥锁同步是比较低级的做法。我们现在改用无缓存的管道来实现同步
	done := make(chan int)
	go func() {
		fmt.Println("Hello, world !")
		<-done
	}()

	done <- 1
}
func chanFunc2() {
	// 上面的代码虽然可以正确同步，但是对管道的缓存大小太敏感：
	// 如果管道有缓存的话，就无法保证main退出之前后台线程能正常打印了。
	// 更好的做法是将管道的发送和接收方向调换一下，这样可以避免同步事件受管道缓存大小的影响：
	done := make(chan int, 1)
	go func() {
		fmt.Println("Hello, world !")
		done <- 1
	}()
	<-done
}
func chanFuncNWork() {
	/*
		对于带缓冲的Channel，对于Channel的第K个接收完成操作发生在第K+C个发送操作完成之前，其中C是Channel的缓存大小。
		虽然管道是带缓存的，main线程接收完成是在后台线程发送开始但还未完成的时刻，此时打印工作也是已经完成的。
	*/
	// 基于带缓存的管道，我们可以很容易将打印线程扩展到N个。下面的例子是开启10个后台线程分别打印：

	// 带 10 个缓存
	done := make(chan int, 10)

	// 开N个后台打印线程
	for i := 0; i < cap(done); i++ {
		go func() {
			fmt.Println("Hello, world !")
			done <- 1
		}()
	}

	// 等待N个后台线程完成
	for i := 0; i < cap(done); i++ {
		<-done
	}

}

func chanFuncNWorkWaitGroup() {
	// 对于这种要等待N个线程完成后再进行下一步的同步操作有一个简单的做法，就是使用sync.WaitGroup来等待一组事件：
	var wg sync.WaitGroup

	// 开N个后台打印线程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("Hello, world !")
			wg.Done()
		}()
	}

	// 等待N个后台线程完成
	wg.Wait()

	/*
		其中wg.Add(1)用于增加等待事件的个数，必须确保在后台线程启动之前执行（如果放到后台线程之中执行则不能保证被正常执行到）。
		当后台线程完成打印工作之后，调用wg.Done()表示完成一个事件。main函数的wg.Wait()是等待全部的事件完成。
	*/
}

func main() {
	//mutexFunc()
	//chanFunc()
	//chanFunc2()
	//chanFuncNWork()
	chanFuncNWorkWaitGroup()
}
