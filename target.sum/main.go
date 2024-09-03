package main

func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, target+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < target+1; j++ {
			if i == 0 {
				if nums[i] == j || (-1*nums[i]) == j {
					dp[i][j] = 1
				}
				continue
			}
			minusCase := dp[i-1][j+nums[i]]

			plusCase := dp[i-1][j-nums[i]]

			dp[i][j] = minusCase + plusCase
		}
	}

	return dp[n-1][target]
}
