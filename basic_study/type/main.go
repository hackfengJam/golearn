package main

import "fmt"

func get(_type interface{})  {
	switch _type.(type) {
	case string:
		print("s")
	case []int:
		print("b")

	}
}
func main() {
	a:= []int{1,2, 3,4}

	fmt.Printf("v1 type:%T\n", a)
	get(a)
}