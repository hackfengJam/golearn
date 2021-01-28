package main

import "fmt"

func main() {
	s := make([]int, 5)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	s = append(s, 1, 2, 3)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println(s)
}
