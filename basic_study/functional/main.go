package main

import (
	"fmt"
	"golearn/basic_study/functional/fib"
)

func main() {
	c := fib.FibnoacciChan()
	for v := range c {
		fmt.Println(v)
	}
}
