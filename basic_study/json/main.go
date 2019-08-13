package main

import (
	"encoding/json"
	"fmt"
)

type Student struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

type Class struct{
	Id int `json:"id"`
	Students []Student `json:"students"`
}

type School struct{
	Id int `json:"id"`
	Classes map[int64]Class `json:"classes"`
}

func basicJSON()  {
	var stu Student

	sJson := `{
	"id": 1,
	"name": "zhangsan",
	"age": 3
}`
	if err := json.Unmarshal([]byte(sJson), &stu); err != nil{
		fmt.Println(err)
	}
	fmt.Println(stu.Id)
}
func arrJSON()  {
	var cls Class

	sJson := `{
	"id": 2,
	"students": [
{
	"id": 1,
	"name": "zhangsan",
	"age": 3
},
{
	"id": 2,
	"name": "lisi",
	"age": 30
},
{
	"id": 3,
	"name": "xiaoming",
	"age": 30
}
]
}`

	if err := json.Unmarshal([]byte(sJson), &cls); err != nil{
		fmt.Println(err)
	}
	//fmt.Println(cls.Id)
	fmt.Println(cls.Students)
}

func mapJSON(){
	var cls School

	sJson := `{
	"id": 2,
	"classes": {
		"1" : {
			"id": 2,
			"students": [
		{
			"id": 1,
			"name": "zhangsan",
			"age": 3
		},
		{
			"id": 2,
			"name": "lisi",
			"age": 30
		},
		{
			"id": 3,
			"name": "xiaoming",
			"age": 30
		}]}, 
		"2" : {
			"id": 2,
			"students": [
		{
			"id": 1,
			"name": "zhangsan2",
			"age": 3
		},
		{
			"id": 2,
			"name": "lisi2",
			"age": 30
		},
		{
			"id": 3,
			"name": "xiaoming2",
			"age": 30
		}
		]}
		}
}`

	if err := json.Unmarshal([]byte(sJson), &cls); err != nil{
		fmt.Println(err)
	}
	//fmt.Println(cls.Id)
	fmt.Println(cls.Classes)
}


func main() {
	//arrJSON()
	mapJSON()
}
