package main

func WordBreak(s string, wordDict []string) bool {
	words := make(map[string]bool)
	for _, v := range wordDict {
		words[v] = true
	}

	n := len(s)

	dp := make([]bool, n+1)
	dp[n] = true

	if words[s[n-1:]] {
		dp[n-1] = true
	}

	for idx := n - 2; idx >= 0; idx-- {
		if words[s[idx:]] {
			dp[idx] = true
		} else {
			for j := 0; idx+j < len(s); j++ {
				dp[idx] = words[s[idx:idx+j+1]] && dp[idx+j+1]
				if dp[idx] {
					break
				}
			}
		}

	}
	return dp[0]
}
