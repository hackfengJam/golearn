package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	// 小写为私有成员，json解码无法访问
	name string `json:"name"`
}

func main() {
	js := `{
		"name":"11"
	}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people: ", p)
}
