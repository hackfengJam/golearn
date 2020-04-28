package init_1

import "fmt"

func init() {
	fmt.Println("in init_1")
}

func Add(a, b int64) int64 {
	return a + b
}
