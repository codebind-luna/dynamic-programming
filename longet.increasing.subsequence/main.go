package main

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n+1; j++ {
			if i == 0 {
				if j == n {
					dp[i][j] = 1
					continue
				}

				if nums[i] < nums[j] {
					dp[i][j] = 1
					continue
				} else {
					dp[i][j] = 0
					continue
				}
			}

			nonpick := 0 + dp[i-1][j]

			if j == n || nums[i] < nums[j] {
				dp[i][j] = max(nonpick, 1+dp[i-1][i])
				continue
			}

			dp[i][j] = nonpick
		}
	}
	return dp[n-1][n]
}
