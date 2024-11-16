package main

func palindrome(i, j int, s string, dp [][]int) int {
	if i >= j {
		return 1
	}

	if s[i] != s[j] {
		return 0
	}

	if dp[i][j] != -1 {
		return dp[i][j]
	}

	dp[i][j] = palindrome(i+1, j-1, s, dp)
	return dp[i][j]
}

func longestPalindrome(s string) string {
	n := len(s)

	if n == 0 || n == 1 {
		return s
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	maxLen := 0
	startIdx := -1

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if palindrome(i, j, s, dp) == 1 {
				if j-i+1 > maxLen {
					maxLen = j - i + 1
					startIdx = i
				}
			}
		}
	}

	return s[startIdx : startIdx+maxLen]
}
