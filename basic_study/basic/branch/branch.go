package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score < 100:
		g = "A"
	default:
		panic(fmt.Sprintf("Wrong "))
	}
	return g
}

func main() {
	const filename = "abc.txt"
	//contents, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%s\n", contents)
	//}

	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}
