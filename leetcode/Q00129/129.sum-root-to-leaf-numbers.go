/*
 * @lc app=leetcode id=129 lang=golang
 *
 * [129] Sum Root to Leaf Numbers
 */
package q00129

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
func sumNumbers(root *TreeNode) int {
	return dfs(root, 0, 0)
}

func dfs(node *TreeNode, result, current int) int {
	if node != nil {
		current = current*10 + node.Val
		if node.Left == nil && node.Right == nil {
			result += current
		}

		return result + dfs(node.Left, result, current) + dfs(node.Right, result, current)
	}

	return 0
}

// @lc code=end
