package main

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = 1
			} else {
				// move upwords
				choiceOne, choiceTwo := 0, 0
				if i > 0 {
					choiceOne = dp[i-1][j]
				}

				if j > 0 {
					choiceTwo = dp[i][j-1] // move leftwords
				}

				dp[i][j] = choiceOne + choiceTwo
			}
		}
	}

	return dp[m-1][n-1]
}
