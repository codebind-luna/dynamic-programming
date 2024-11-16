package main

import (
	"strconv"
)

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	n := len(s)
	dp := make([]int, n+1)
	dp[n] = 1

	if s[n-1] != '0' {
		dp[n-1] = 1
	}

	for i := n - 2; i >= 0; i-- {
		if s[i] == '0' {
			dp[i] = 0
		} else {
			for j := 0; j < 2; j++ {
				val, _ := strconv.Atoi(s[i : i+j+1])
				if 1 <= val && val <= 26 {
					dp[i] += dp[i+j+1]
				}
			}
		}
	}

	return dp[0]
}
