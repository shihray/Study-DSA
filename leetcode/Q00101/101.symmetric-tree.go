/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
 */
package q00101

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
func isSymmetric(root *TreeNode) bool {
	return isSame(root, root)
}

func isSame(left, right *TreeNode) bool {
	// left := root.Left
	// right := root.Right

	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}
	return isSame(left.Right, right.Left) && isSame(left.Left, right.Right)
}

// @lc code=end
