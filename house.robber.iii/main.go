package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode) [2]int {
	if root == nil {
		return [2]int{0, 0}
	}

	left := dfs(root.Left)
	right := dfs(root.Right)

	withRoot := root.Val + left[1] + right[1]
	withOutRoot := max(left[0], left[1]) + max(right[0], right[1])
	return [2]int{withRoot, withOutRoot}
}

func rob(root *TreeNode) int {
	result := dfs(root)
	maxWithRoot, maxWithoutRoot := result[0], result[1]

	return max(maxWithRoot, maxWithoutRoot)
}
