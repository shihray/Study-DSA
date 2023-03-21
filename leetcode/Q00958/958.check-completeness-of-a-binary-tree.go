/*
 * @lc app=leetcode id=958 lang=golang
 *
 * [958] Check Completeness of a Binary Tree
 */
package q00958

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
func isCompleteTree(root *TreeNode) bool {
	return dfs(root, 0, countNodes(root))
}

func dfs(node *TreeNode, currentIndex, nodes int) bool {
	if node == nil {
		return true
	}

	if currentIndex >= nodes {
		return false
	}

	return dfs(node.Left, 2*currentIndex+1, nodes) && dfs(node.Right, 2*currentIndex+2, nodes)
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

// @lc code=end
