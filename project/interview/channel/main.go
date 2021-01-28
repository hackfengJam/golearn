package main

import "fmt"

type query func(string) string

func exec(name string, vs ...query) string {
	ch := make(chan string)
	fn := func(i int) {
		ch <- vs[i](name)
	}
	for i, _ := range vs {
		go fn(i)
	}
	return <-ch
}

func main() {
	ret := exec("111", func(n string) string {
		return n + "func1"
	}, func(n string) string {
		return n + "func2"
	}, func(n string) string {
		return n + "func3"
	}, func(n string) string {
		return n + "func4"
	})
	/*
		依据4个goroutine的启动后执行效率，很可能打印111func4，
		但其他的111func*也可能先执行，exec只会返回一条信息

		因为 exec中<-ch只执行一次，所以不管多少 query 传进去，只返回最先执行完的那个.
		所以这个exec是有问题的，会导致goroutine泄漏
	*/
	fmt.Println(ret)
}
