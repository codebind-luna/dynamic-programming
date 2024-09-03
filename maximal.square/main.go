package main

import (
	"fmt"
	"sort"
)

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])

	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	maxSide := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
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
				dp[i][j] = min(append([]int{}, up, diag, left)) + 1
				maxSide = max(maxSide, dp[i][j])
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", dp[i][j])
		}
		fmt.Println()
	}

	return maxSide * maxSide
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(arr []int) int {
	sort.Ints(arr)
	return arr[0]
}

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
	}

	maximalSquare(matrix)
}
