package main

import (
	"fmt"
	"strings"
)

func main() {
	//var client *api.Client
	//var m map[int]int
	//m = make(map[int]int, 10)
	//m[1] = 1
	//m[2] = 2
	//m[3] = 3
	//
	//for k, v := range m {
	//	fmt.Printf("k: %v -> v: %v\n", k, v)
	//}

	str := " 123 123 "
	str = strings.TrimSpace(str)
	fmt.Printf("---%v---", str)

}
