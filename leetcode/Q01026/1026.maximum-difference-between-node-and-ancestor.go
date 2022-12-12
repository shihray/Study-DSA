/*
 * @lc app=leetcode id=1026 lang=golang
 *
 * [1026] Maximum Difference Between Node and Ancestor
 */
package q01026

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
func maxAncestorDiff(root *TreeNode) int {
    return dfs(root, root.Val, root.Val)
}

func dfs(root *TreeNode, maxVal, minVal int) int {
	if root == nil {
		return maxVal - minVal
	}
	maxVal = max(maxVal, root.Val)
	minVal = min(minVal, root.Val)

	return max(dfs(root.Left, maxVal, minVal), dfs(root.Right, maxVal, minVal))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

