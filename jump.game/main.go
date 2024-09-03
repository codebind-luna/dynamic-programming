package main

func canJump(nums []int) bool {
	n := len(nums)

	dp := make([]bool, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i] = (nums[i] == 0)
			continue
		}

		ans := false
		for j := nums[i]; j > 0; j-- {
			ans = ans || dp[i-j]
		}
		return ans
	}
	return dp[n-1]
}
