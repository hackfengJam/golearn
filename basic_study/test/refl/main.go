package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func arrayT() {
	slice := []string{"hello", "world"}
	type sliceHeader struct {
		Data unsafe.Pointer
		Len  int
		Cap  int
	}
	header := (*sliceHeader)(unsafe.Pointer(&slice))
	fmt.Println(header.Len)
	elementType := reflect.TypeOf(slice).Elem()
	fmt.Println(header.Data)
	fmt.Println(elementType.Size())
	secondElementPtr := uintptr(header.Data) + elementType.Size()
	*((*string)(unsafe.Pointer(secondElementPtr))) = "!!!"
	fmt.Println(slice)
}

type AA struct {
	x *BB
	y int64
}

type BB struct {
	x int64
	y int64
}

func main() {
	arrayT()

	// a := AA{x: &BB{x: 1, y: 2}, y: 2}
	// elementType := reflect.TypeOf(a).Elem()
	// fmt.Println(elementType)

}
