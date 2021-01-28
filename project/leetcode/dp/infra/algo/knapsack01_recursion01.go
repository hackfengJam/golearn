package algo

import (
	"math"
)

func bestValueRecursion(w, v []int, index int, c int) int {
	if index < 0 || c <= 0 {
		return 0
	}
	var res = bestValueRecursion(w, v, index-1, c)
	if c >= w[index] {
		res = int(math.Max(float64(res), float64(bestValueRecursion(w, v, index-1, c-w[index])+v[index])))
	}

	return res
}

/*
f(i, c) = f(i-1, c)
        = f(i-1, c-w[i]) + v[i]

*/
func Knapsack01Recursion(w, v []int, C int) int {
	return bestValueRecursion(w, v, len(w)-1, C)
}
