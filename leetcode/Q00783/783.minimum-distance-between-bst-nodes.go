/*
 * @lc app=leetcode id=783 lang=golang
 *
 * [783] Minimum Distance Between BST Nodes
 */
package q00783

import "math"

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
func minDiffInBST(root *TreeNode) int {
	min := math.MaxInt32
	prevVal := math.MinInt32
	helper := func(node *TreeNode) {}
	helper = func(n *TreeNode) {
		if n == nil {
			return
		}

		helper(n.Left)
		if prevVal != math.MinInt32 {
			d := n.Val - prevVal
			if d < min {
				min = d
			}
		}

		prevVal = n.Val
		helper(n.Right)
	}

	helper(root)

	return min
}

// @lc code=end
