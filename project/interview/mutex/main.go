package main

import "sync"

/*
https://yushuangqi.com/blog/2017/golang-mian-shi-ti-da-an-yujie-xi.html?from=groupmessage&isappinstalled=0

虽然有使用sync.Mutex做写锁，但是map是并发读写不安全的。
map属于引用类型，并发读写时多个协程见是通过指针访问同一个地址，即访问共享变量，此时同时读写资源存在竞争关系。
会报错误信息:“fatal error: concurrent map read and map write”。

可以在在线运行中执行，复现该问题。那么如何改善呢? 当然Go1.9新版本中将提供并发安全的map。首先需要了解两种锁的不同：

sync.Mutex互斥锁
sync.RWMutex读写锁，基于互斥锁的实现，可以加多个读锁或者一个写锁。
利用读写锁可实现对map的安全访问，在线运行改进版 。利用 RWutex 进行读锁。

//author: ysqi ,https://yushuangqi.com

package main

import (
	"fmt"
	"sync"
)

type UserAges struct {
	ages map[string]int
	sync.RWMutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	ua.RLock()
	defer ua.RUnlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}

	return -1
}

func main() {
	count := 10000
	gw := sync.WaitGroup{}
	gw.Add(count * 3)
	u := UserAges{ages: map[string]int{}}
	add := func(i int) {
		u.Add(fmt.Sprintf("user_%d", i), i)
		gw.Done()
	}
	for i := 0; i < count; i++ {
		go add(i)
		go add(i)
	}
	for i := 0; i < count; i++ {
		go func(i int) {
			defer gw.Done()
			u.Get(fmt.Sprintf("user_%d", i))
			fmt.Print(".")
		}(i)
	}
	gw.Wait()
	fmt.Println("Done")
}


*/
type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}
