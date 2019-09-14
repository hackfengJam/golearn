package algo

import (
	"math"
)

var memo [][]int

func bestValueRecursion02(w, v []int, index int, c int) int {
	if index < 0 || c <= 0 {
		return 0
	}
	if memo[index][c] != -1 {
		return memo[index][c]
	}

	var res = bestValueRecursion(w, v, index-1, c)
	if c >= w[index] {
		res = int(math.Max(float64(res), float64(bestValueRecursion(w, v, index-1, c-w[index])+v[index])))
	}

	memo[index][c] = res

	return memo[index][c]
}

/*
f(i, c) = f(i-1, c)
        = f(i-1, c-w[i]) + v[i]

*/
func Knapsack01Recursion02(w, v []int, C int) int {

	// init memo
	memo = make([][]int, len(w))
	for i := 0; i < len(w); i++ {
		memo[i] = make([]int, C+1)
		for j := 0; j <= C; j++ {
			memo[i][j] = -1
		}
	}

	return bestValueRecursion02(w, v, len(w)-1, C)
}
