package main

func numSubseq(nums []int, target int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, target+1)
	}

	for i := 0; i < n; i++ {
		dp[i][0] = 1
	}

	if nums[0] <= target {
		dp[0][nums[0]] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < target+1; j++ {
			if nums[i] <= j {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}

		}
	}

	return dp[n-1][target]
}
