package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Hello goroutine!");
	}()
	fmt.Println("Hello world!")
	time.Sleep(time.Millisecond)
}
