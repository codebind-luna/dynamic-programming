package main

import "log"

func helper(result *[][]int, output []int, nums []int, start int) {
	if start == len(nums) {
		*result = append(*result, append([]int{}, output...))
		return
	}

	output = append(output, nums[start])
	helper(result, output, nums, start+1)
	output = output[:len(output)-1]
	helper(result, output, nums, start+1)
}

func main() {
	nums := []int{3, 2, 1}

	var result [][]int
	helper(&result, []int{}, nums, 0)
	log.Println(result)
}
