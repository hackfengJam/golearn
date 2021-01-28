// +build wechat_debug

package main

import "fmt"

func returnMap() map[string]string {
	return nil
}

func Add() {
	fmt.Println(returnMap()["sign"])
	return
}
func main() {
	Add()
}
