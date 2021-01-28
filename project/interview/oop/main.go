package main

import (
	"fmt"
	"time"
)

type Project struct{}

func (p *Project) deferError() {
	if err := recover(); err != nil {
		fmt.Println("recover: ", err)
	}
}

func (p *Project) exec(msgchan chan interface{}) {
	for msg := range msgchan {
		m := msg.(int)
		fmt.Println("msg: ", m)
		//if m, ok := msg.(int); ok {
		//	fmt.Println("msg: ", m)
		//}
	}
}

func (p *Project) run(msgchan chan interface{}) {
	for {
		defer p.deferError()
		go p.exec(msgchan)
		time.Sleep(time.Second * 2)
	}
}

func (p *Project) Main() {
	a := make(chan interface{}, 100)
	go p.run(a)
	go func() {
		for {
			a <- "1"
			//a <- 1
			time.Sleep(time.Second)
		}
	}()

	// constant 100000000000000000000000 overflows time.Duration
	time.Sleep(time.Second * 1000)
	//time.Sleep(time.Second * 100000000000000)
}

func main() {
	/*
		new：https://golang.org/pkg/builtin/#new
		内置函数 new 分配空间。传递给new 函数的是一个类型，不是一个值。
		返回值是 指向这个新分配的零值的指针。

		make：https://golang.org/pkg/builtin/#make
		内建函数 make 分配并且初始化 一个 slice, 或者 map 或者 chan 对象。
		并且只能是这三种对象。	和 new 一样，第一个参数是 类型，不是一个值。
		但是make 的返回值就是这个类型（即使一个引用类型），而不是指针。 具体的返回值，依赖具体传入的类型。
	*/
	p := new(Project)
	p.Main()
}
