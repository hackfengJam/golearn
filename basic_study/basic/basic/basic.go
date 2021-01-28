package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"reflect"
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

	type Label int64
	const (
		_ Label = iota
		allReviewsLabel
		allReviewsLabel2
		allReviewsLabel3
		allReviewsLabel4
		allReviewsLabel5
		allReviewsLabel6
	)

	fmt.Println(allReviewsLabel, allReviewsLabel2, allReviewsLabel3, allReviewsLabel4, allReviewsLabel5, allReviewsLabel6)
	fmt.Println(reflect.TypeOf(allReviewsLabel), reflect.ValueOf(reflect.TypeOf(allReviewsLabel)))

	s := []byte{123, 34, 110, 97, 109, 101, 115, 112, 97, 99, 101, 34, 58, 34, 97, 114, 101, 110, 97, 58, 97, 112, 105, 58, 101, 118, 97, 108, 117, 97, 116, 105, 111, 110, 58, 100, 97, 110, 109, 117, 34, 34, 44, 34, 117, 105, 100, 34, 58, 49, 48, 48, 48, 48, 51, 44, 34, 110, 105, 99, 107, 110, 97, 109, 101, 34, 58, 34, 229, 176, 143, 232, 144, 140, 230, 142, 168, 232, 141, 144, 34, 44, 34, 97, 118, 97, 116, 97, 114, 34, 58, 34, 111, 115, 115, 47, 49, 47, 50, 48, 49, 56, 48, 55, 47, 102, 49, 101, 48, 53, 100, 99, 98, 97, 51, 100, 101, 51, 54, 55, 98, 45, 51, 55, 120, 51, 54, 46, 112, 110, 103, 34, 125}
	fmt.Println(string(s))
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
	//fmt.Println(strings.Replace("123 34 110 97 109 101 115 112 97 99 101 34 58 34 97 114 101 110 97 58 97 112 105 58 101 118 97 108 117 97 116 105 111 110 58 100 97 110 109 117 34 44 34 114 101 119 97 114 100 34 58 34 229 136 154 229 136 154 229 133 165 233 128 137 228 186 134 231 178 190 233 128 137 230 181 139 232 175 132 34 44 34 117 105 100 34 58 49 48 48 48 48 51 44 34 110 105 99 107 110 97 109 101 34 58 34 229 176 143 232 144 140 230 142 168 232 141 144 34 44 34 97 118 97 116 97 114 34 58 34 111 115 115 47 49 47 50 48 49 56 48 55 47 102 49 101 48 53 100 99 98 97 51 100 101 51 54 55 98 45 51 55 120 51 54 46 112 110 103 34 125", " ", ",", -1))
}
