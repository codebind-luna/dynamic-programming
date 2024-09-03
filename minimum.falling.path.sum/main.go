package main

func minFallingPathSum(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])

	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	copy(dp[0], matrix[0])

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			var minPathAbove int

			switch j {
			case 0:
				minPathAbove = min([]int{dp[i-1][j], dp[i-1][j+1]})
			case len(dp[0]) - 1:
				minPathAbove = min([]int{dp[i-1][j-1], dp[i-1][j]})
			default:
				minPathAbove = min([]int{dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1]})
			}

			dp[i][j] = minPathAbove + matrix[i][j]

		}
	}
	return min(dp[len(dp)-1])
}

func min(a []int) int {
	curMin := a[0]
	for _, elem := range a {
		if elem < curMin {
			curMin = elem
		}
	}
	return curMin
}

func main() {
	matrix := [][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}

	minFallingPathSum(matrix)
}
