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

	if nums[0] == 0 {
		dp[0][0] = 2
	}

	for i := 1; i < n; i++ {
		for j := 1; j < target+1; j++ {
			nonPick := dp[i-1][j]

			pick := 0
			if nums[i] <= j {
				pick = dp[i-1][j-nums[i]]
			}

			dp[i][j] = nonPick + pick

		}
	}

	return dp[n-1][target]
}
