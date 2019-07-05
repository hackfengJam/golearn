package main

import "fmt"

type Student struct {
	name string
}

func main() {
	/*

		map 中 value 是非指针类型，可以当做一个原子变量，只能统一更改，不能只改 Student 其中一个成员，
		如果想要更改 Student 其中一个成员，使用 map[string]*Student
	*/
	//m := map[string]Student{"people": {"zhoujielun"}}
	m := map[string]*Student{"people": {"zhoujielun"}}
	m["people"].name = "wuyanzu"
	fmt.Println(m)
}
