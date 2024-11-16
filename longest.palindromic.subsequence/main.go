package main

func reverse(s string) string {
	s1 := ""

	for i := 0; i < len(s); i++ {
		s1 = string(s[i]) + s1
	}
	return s1
}

func longestPalindromeSubseq(s string) int {
	text1 := s
	text2 := reverse(s)
	n := len(text1)
	m := len(text2)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
				continue
			}

			dp[i][j] = max(dp[i+1][j], dp[i][j+1])
		}
	}

	return dp[0][0]
}
