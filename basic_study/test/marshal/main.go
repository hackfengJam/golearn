package main

import (
	"encoding/json"
	"fmt"
)

type Chan struct {
	A *Chan  `json:"a"`
	B int    `json:"b"`
	C string `json:"c"`
	//D chan int `json:"d"` // unsupported type: chan int
}

func main() {
	c := &Chan{
		A: &Chan{
		},
		B: 1,
		C: "c",
		//D: make(chan int, 0), // unsupported type: chan int
	}
	var x, y *Chan
	x = c
	for i := 0; i < 100000; i++ {
		y = x.A
		y.A = &Chan{}
		x = y
	}

	s, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))

}
