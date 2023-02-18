/*
 * @lc app=leetcode id=226 lang=golang
 *
 * [226] Invert Binary Tree
 */
package q00226

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
func invertTree(root *TreeNode) *TreeNode {
	if root != nil {

		root.Left, root.Right = root.Right, root.Left
		invertTree(root.Left)
		invertTree(root.Right)

		return root

	} else {
		return nil
	}
}

// @lc code=end
