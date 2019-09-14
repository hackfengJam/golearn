package algo

import (
	"math"
)

/*
f(i, c) = f(i-1, c)
        = f(i-1, c-w[i]) + v[i]

*/
func Knapsack01Dp02(w, v []int, C int) int {
	var memo [][]int
	// init memo
	memo = make([][]int, 2)
	for i := 0; i < 2; i++ {
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
			memo[i%2][j] = memo[(i-1)%2][j]
			if j >= w[i] {
				memo[i%2][j] = int(math.Max(float64(memo[i%2][j]), float64(memo[(i-1)%2][j-w[i]]+v[i])))
			}
		}
	}

	return memo[(len(w)-1)%2][C]
}
