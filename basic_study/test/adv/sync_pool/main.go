package main

import (
	"fmt"
	"sync"
)

func A() {

}

func main() {
	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}

	a := p.Get().(int)
	p.Put(1)
	p.Put("aa")
	b := p.Get().(int)
	c := p.Get().(string)
	d := p.Get().(int)
	fmt.Println(a, b, c, d)
}
