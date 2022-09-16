/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 */

package q00094

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
func inorderTraversal(root *TreeNode) []int {
    var result []int
	if root == nil {
		return result
	}
	Dfs(root, &result)
	return result
}

func Dfs(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	Dfs(node.Left, res)
	(*res) = append((*res), node.Val)
	Dfs(node.Right, res)
}
// @lc code=end

