package main

import "math"

// idx == 0
// return 0
// notPick := f(idx-1, prev, )
// pick := 1 + f(idx-1, s[idx])

func maxLen(idx int, last byte, k int, s string, dp [][]int) int {
	if idx == 0 {
		if last == '{' || abs(int(s[idx]-last)) <= k {
			return 1
		}
		return 0
	}

	if dp[idx][last-'a'] != -1 {
		return dp[idx][last-'a']
	}

	notPick := maxLen(idx-1, last, k, s, dp)

	pick := math.MinInt
	if last == ' ' || abs(int(s[idx]-last)) <= k {
		pick = 1 + maxLen(idx-1, s[idx], k, s, dp)
	}

	dp[idx][last-'a'] = max(notPick, pick)

	return dp[idx][last-'a']
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func longestIdealString(s string, k int) int {
	n := len(s)
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, 66)
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= 26; j++ {
			notPick := dp[i-1][j]
			pick := math.MinInt
			if j == 26 || abs(int(s[i]-('a'+byte(j)))) <= k {
				pick = 1 + dp[i-1][s[i]-'a']
			}
			dp[i][j] = max(notPick, pick)
		}
	}

	return dp[n-1][26]
}
