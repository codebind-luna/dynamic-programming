package main

func helper(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	dp[0] = nums[0]

	for i := 1; i < n; i++ {
		take := nums[i]

		if i > 1 {
			take += dp[i-2]
		}
		notTake := dp[i-1]

		dp[i] = max(notTake, take)
	}

	return dp[n-1]
}

func rob(nums []int) int {
	numsWithoutFirst := nums[1:]
	numsWithoutLast := nums[:len(nums)-1]

	return max(helper(numsWithoutFirst), helper(numsWithoutLast))
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
