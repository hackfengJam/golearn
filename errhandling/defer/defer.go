package main

import (
	"bufio"
	"fmt"
	"learngo/functional/fib"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occurred")
	fmt.Println(4)
}

func tryDefer2() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	f := fib.Fibnoacci()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}

}

func writeFile2(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	f := fib.Fibnoacci()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}

}

func main() {
	tryDefer2()
	//writeFile("fib.txt")
}
