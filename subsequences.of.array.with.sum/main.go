package main

import "log"

func helper(result *[][]int, output []int, nums []int, start int, sum int, target int) {
	if start == len(nums) {
		if sum == target {
			*result = append(*result, append([]int{}, output...))
		}
		return
	}

	output = append(output, nums[start])
	sum += nums[start]
	helper(result, output, nums, start+1, sum, target)
	output = output[:len(output)-1]
	sum -= nums[start]
	helper(result, output, nums, start+1, sum, target)
}

// Driver Code
func main() {
	arr := []int{1, 2, 3, 4, 5}
	var result [][]int

	helper(&result, []int{}, arr, 0, 0, 5)

	log.Println(result)
}
