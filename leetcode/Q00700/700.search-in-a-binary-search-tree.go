/*
 * @lc app=leetcode id=700 lang=golang
 *
 * [700] Search in a Binary Search Tree
 */
package q00700

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
func searchBST(root *TreeNode, val int) *TreeNode {

	if root == nil {
		return nil
	}

	if val == root.Val {
		return root
	} else if val > root.Val {
		return searchBST(root.Right, val)
	}
	return searchBST(root.Left, val)
}
// @lc code=end

