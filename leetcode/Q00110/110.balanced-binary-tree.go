/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
 */
package q00110

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
func isBalanced(root *TreeNode) bool {
	_, ok := dfs(root)
	return ok
}

func dfs(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	l, ok := dfs(root.Left)
	if !ok {
		return 0, false
	}
	r, ok := dfs(root.Right)
	if !ok {
		return 0, false
	}

	if diff := abs(l - r) ; diff > 1 {
		return 0, false
	}
	return max(l, r) + 1, true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end
