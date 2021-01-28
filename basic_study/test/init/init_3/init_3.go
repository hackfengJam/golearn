package init_3

import (
	"fmt"
	"golearn/basic_study/test/init/init_1"
)

func init() {
	fmt.Println("in init_3")
}

func Add(a, b int64) int64 {
	return init_1.Add(a, b)
}
