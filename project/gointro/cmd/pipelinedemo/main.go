package main

import (
	"../../pipeline"
	"bufio"
	"fmt"
	"os"
)

func main() {
	//const filename = "large.in"
	const filename = "small.in"
	//const n = 100000000
	const n = 64
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.RandomSource(n)
	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)
	writer.Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}

	count := 0
	defer file.Close()
	p = pipeline.ReaderSource(bufio.NewReader(file), -1)

	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}

func mergeDemo() {
	p := pipeline.Merge(
		pipeline.InMemSort(pipeline.ArraySource(3, 2, 6, 7, 4)),
		pipeline.InMemSort(pipeline.ArraySource(7, 4, 0, 3, 2, 8, 13, 8)))
	for v := range p {
		fmt.Println(v)
	}
}
