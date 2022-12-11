/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
 */
package q00124

import "math"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    maxVal := math.MinInt64
	dfs(root, &maxVal)
	return maxVal
}

func dfs(root *TreeNode, maxVal *int) int {
	if root == nil {
		return 0
	}

	left := max( dfs(root.Left, maxVal), 0)
	right := max( dfs(root.Right, maxVal), 0)

	*maxVal = max(*maxVal, root.Val + left + right)
	return root.Val + max(left ,right)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
// @lc code=end

