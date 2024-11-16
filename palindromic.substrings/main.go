package main

func countSubstrings(s string) int {
	n := len(s)

	if n == 0 || n == 1 {
		return 1
	}

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	count := 0

	for i := n - 1; i >= 0; i-- {
		j := n - 1
		for k := 0; k < n-i; k++ {
			// all single char substrings are palindromic
			if i == j {
				dp[i][j] = true
				count++
				j--
				continue
			}

			// check for two char substrings if palindromic
			if j == i+1 {
				if s[i] == s[j] {
					dp[i][j] = true
					count++
				}
				j--
				continue
			}

			// check for more than two char substrings if palindromic
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				count++
			}
			j--
		}
	}

	return count
}
