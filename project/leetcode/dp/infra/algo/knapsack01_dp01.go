package algo

import (
	"fmt"
	"math"
)

/*
f(i, c) = f(i-1, c)
        = f(i-1, c-w[i]) + v[i]

*/
func Knapsack01Dp01(w, v []int, C int) int {
	var memo [][]int
	// init memo
	memo = make([][]int, len(w))
	for i := 0; i < len(w); i++ {
		memo[i] = make([]int, C+1)
		for j := 0; j <= C; j++ {
			memo[i][j] = -1
		}
	}

	// init start state
	for j := 0; j <= C; j++ {
		memo[0][j] = 0
		if j >= w[0] {
			memo[0][j] = v[0]
		}
	}

	// dp
	for i := 1; i < len(w); i++ {
		for j := 0; j <= C; j++ {
			memo[i][j] = memo[i-1][j]
			if j >= w[i] {
				memo[i][j] = int(math.Max(float64(memo[i][j]), float64(memo[i-1][j-w[i]]+v[i])))
			}
		}
	}

	// backtrack
	var ret []int
	i := len(w) - 1
	j := C
	for {
		if j == 0 {
			break
		}
		//if i == 0 {
		//	break
		//}
		if memo[i][j] != memo[i-1][j] {
			ret = append(ret, i)
			j -= w[i]
		}
		i--
	}
	fmt.Println(ret)

	return memo[len(w)-1][C]
}
