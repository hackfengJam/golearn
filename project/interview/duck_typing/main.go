package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Stduent struct{}

// 编译失败，值类型 Student{} 未实现接口People的方法，不能定义为 People类型。
//func (stu *Stduent) Speak(think string) (talk string) {
//	if think == "bitch" {
//		talk = "You are a good boy"
//	} else {
//		talk = "hi"
//	}
//	return
//}

// 编译通过
func (stu Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo People = Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
