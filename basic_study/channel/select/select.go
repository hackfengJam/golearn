package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, c chan int) {
	//for {
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("Worker %d received %c\n", id, n)
	//}
	time.Sleep(time.Second)
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}
func main() {
	var c1, c2 = generator(), generator()
	//w := createWorker(0)
	var worker = createWorker(0)

	var values [] int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	n := 0
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
		//default:
		//	fmt.Println("No value received")
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue len=", len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}
