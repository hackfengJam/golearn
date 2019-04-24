package main

import "fmt"

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(r)
		}
	}()

	//a := 0
	//a = 5 / a
	panic(123)

}

func main() {
	tryRecover()
}
