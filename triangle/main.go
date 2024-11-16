package triangle

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for j := 0; j < n; j++ {
		dp[n-1][j] = triangle[n-1][j]
	}

	for i := n - 2; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			d := triangle[i][j] + dp[i+1][j]
			dg := triangle[i][j] + dp[i+1][j+1]

			dp[i][j] = min(d, dg)
		}
	}

	return dp[0][0]
}
