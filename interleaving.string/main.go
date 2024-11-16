package main

func canMake(i, j int, s1, s2, s3 string) bool {
	if i == len(s1) && j == len(s2) {
		return true
	}

	if i < len(s1) && s1[i] == s3[i+j] && canMake(i+1, j, s1, s2, s3) {
		return true
	}

	if j < len(s2) && s2[j] == s3[i+j] && canMake(i, j+1, s1, s2, s3) {
		return true
	}

	return false
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	n := len(s1)
	m := len(s2)

	dp := make([][]bool, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}

	for i := n; i >= 0; i-- {
		for j := m; j >= 0; j-- {
			if i == n && j == m {
				dp[i][j] = true
				continue
			}

			if i < n && s1[i] == s3[i+j] && dp[i+1][j] {
				dp[i][j] = true
				continue
			}

			if j < m && s2[j] == s3[i+j] && dp[i][j+1] {
				dp[i][j] = true
				continue
			}
		}
	}

	return dp[0][0]
}
