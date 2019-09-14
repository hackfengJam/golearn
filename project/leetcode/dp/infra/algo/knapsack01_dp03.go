package algo

import (
	"math"
)

/*
f(i, c) = f(i-1, c)
        = f(i-1, c-w[i]) + v[i]

*/
func Knapsack01Dp03(w, v []int, C int) int {
	var memo []int
	// init memo
	memo = make([]int, C+1)
	for i := 0; i <= C; i++ {
		memo[i] = -1
	}

	// init start state
	for i := 0; i <= C; i++ {
		memo[i] = 0
		if i >= w[0] {
			memo[i] = v[0]
		}
	}

	// dp
	for i := 1; i < len(w); i++ {
		for j := C; j >= 0; j-- {
			memo[j] = memo[j]
			if j >= w[i] {
				memo[j] = int(math.Max(float64(memo[j]), float64(memo[j-w[i]]+v[i])))
			}
		}
	}

	return memo[C]
}
