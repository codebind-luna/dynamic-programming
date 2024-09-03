package main

import (
	"log"
	"sort"
)

func countSquares(matrix [][]int) int {

	m := len(matrix)
	n := len(matrix[0])

	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				if i == 0 || j == 0 {
					dp[i][j] = 1
					ans += dp[i][j]
					continue
				}
				left, up, diag := 0, 0, 0
				if i > 0 {
					left = dp[i-1][j]
				}
				if j > 0 {
					up = dp[i][j-1]
				}
				if i > 0 && j > 0 {
					diag = dp[i-1][j-1]
				}
				dp[i][j] = min(append([]int{}, left, up, diag)) + 1
				ans += dp[i][j]
			}
		}
	}

	return ans
}

func min(arr []int) int {
	sort.Ints(arr)
	return arr[0]
}

func main() {
	matrix := [][]int{
		{0, 1, 1, 1},
		{1, 1, 1, 1},
		{0, 1, 1, 1},
	}
	log.Println(countSquares(matrix))
}
