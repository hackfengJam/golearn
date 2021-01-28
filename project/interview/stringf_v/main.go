package main

import "fmt"

type People struct {
	Name string
}

func (p *People) String() string {
	return fmt.Sprintf("print: %v", p)
}

func main() {
	/*
		https://golang.org/pkg/fmt/
		The default format for %v is:
			bool:                    %t
			int, int8 etc.:          %d
			uint, uint8 etc.:        %d, %#x if printed with %#v
			float32, complex64, etc: %g
			string:                  %s
			chan:                    %p
			pointer:                 %p

		// 因此会调用 String 方法，从而 无限递归。
	*/
	p := &People{}
	p.String()
}
