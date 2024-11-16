package main

func LongestPalindrome(s string) string {
	n := len(s)

	if n == 0 || n == 1 {
		return s
	}

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	startIdx := -1
	maxLen := 0

	for i := n - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if i > j {
				dp[i][j] = false
				continue
			}
			// all single char substrings are palindromic
			if i == j {
				dp[i][j] = true
				if maxLen < j-i+1 {
					maxLen = j - i + 1
					startIdx = i
				}
				continue
			}

			// check for two char substrings if palindromic
			if j == i+1 {
				if s[i] == s[j] {
					dp[i][j] = true
					if maxLen < j-i+1 {
						maxLen = j - i + 1
						startIdx = i
					}
				}
				continue
			}

			// check for more than two char substrings if palindromic
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				if maxLen < j-i+1 {
					maxLen = j - i + 1
					startIdx = i
				}
			}
		}
	}

	return s[startIdx : startIdx+maxLen]
}
