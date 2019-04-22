package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//var aa = 3
//var ss = "kkk"
//// bb := true 不可以
//var bb = true
var (
	aa = 3
	ss = "kkkk"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}
func variableInitalValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}
func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}
func euler() {
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c)) // 5

	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)

	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)

	fmt.Printf("%.3f", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
func consts() {
	const filename = "1.txt"
	const (
		a, b     = 3, 4
		i, j int = 3, 4
	)

	var c int
	c = int(math.Sqrt(a*a + b*b))

	var k int
	k = int(math.Sqrt(float64(i*i + j*j)))
	fmt.Println(filename, c, k)
}

func enums() {
	//const (
	//	cpp    = 0
	//	java   = 1
	//	python = 2
	//	golang = 3
	//)
	//const (
	//	cpp = iota // 自增值
	//	java
	//	python
	//	golang
	//)
	const (
		cpp = iota // 自增值
		_          // 中间某个值不需要，使用 _
		java
		python
		golang
	)

	// b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	//fmt.Println("Hello world")
	//variableZeroValue()
	//variableInitalValue()
	//variableTypeDeduction()
	//variableShorter()
	//fmt.Println(aa, ss, bb)
	//
	//euler()
	//triangle()
	//consts()
	enums()
}
