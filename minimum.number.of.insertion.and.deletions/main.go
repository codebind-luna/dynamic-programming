package main

func canYouMake(s1, s2 string) int {
	// Write your code here.
	m := len(s1)
	n := len(s2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Initialize the dp array for the base cases
	for i := 0; i <= m; i++ {
		dp[i][n] = m - i // Remaining characters in word1 to delete
	}
	for j := 0; j <= n; j++ {
		dp[m][j] = n - j // Remaining characters in word2 to insert
	}

	dp[m][n] = 0

	// Fill the dp array from the back
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s1[i] == s2[j] {
				dp[i][j] = dp[i+1][j+1] // No cost if the characters match
			} else {
				choiceInsert := 1 + dp[i][j+1] // Insert
				choiceDelete := 1 + dp[i+1][j] // Delete
				dp[i][j] = min(choiceInsert, choiceDelete)
			}
		}
	}

	return dp[0][0]
}
