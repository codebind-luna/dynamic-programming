package main

import "fmt"

func numSubseq(nums []int, target int) int {
	mod := 1000_000_007
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, target+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < target+1; j++ {
			if i == 0 {
				if j == 0 && nums[0] == 0 {
					dp[0][j] = 2
				} else if j == 0 || nums[0] == j {
					dp[0][j] = 1
				}
			} else {
				notPick := dp[i-1][j]

				pick := 0

				if nums[i] <= j {
					pick = dp[i-1][j-nums[i]]
				}

				dp[i][j] = (notPick + pick) % mod
			}
		}
	}

	return dp[n-1][target]
}

func main() {
	nums := []int{0, 0, 1}
	target := 1

	fmt.Println(numSubseq(nums, target))
}
