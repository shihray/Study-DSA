/*
 * @lc app=leetcode id=1372 lang=golang
 *
 * [1372] Longest ZigZag Path in a Binary Tree
 */
package q01372

type TreeNode struct {
	Val   int
	Left  *TreeNode
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
func longestZigZag(root *TreeNode) int {
	maxVal := 0

	dfs(root, &maxVal)

	return maxVal
}

func dfs(node *TreeNode, maxVal *int) (int, int) {
	if node == nil {
		return -1, -1
	}

	_, val1 := dfs(node.Left, maxVal)
	goLeft := val1 + 1

	val2, _ := dfs(node.Right, maxVal)
	goRight := val2 + 1

	*maxVal = max(*maxVal, max(goLeft, goRight))

	return goLeft, goRight
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
