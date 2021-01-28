package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:

		// panic 是中断执行
		panic("unsupported operation: " + op)
	}
}

// 13 / 3 = 4 ... 1
func div(a, b int) (q, r int) {
	// 实现带余除法
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]

	}
	return s
}

func swap(a, b *int) {
	// *b, *a = *b, *a
	*b, *a = *a, *b
}

func main() {
	//fmt.Println(eval(3, 4, "*"))
	//fmt.Println(div(13, 3))

	//fmt.Println(apply(pow, 3, 4))
	//
	//fmt.Println(apply(func(a, b int) int {
	//	return int(math.Pow(float64(a), float64(b)))
	//}, 3, 4))
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)

}
