/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
 */
package q00144

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
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var resp []int
	bfs(root, &resp)
	return resp
}

func bfs(cur *TreeNode, ans *[]int) {
	if cur == nil {
		return
	}

	*ans = append(*ans, cur.Val)
	bfs(cur.Left, ans)
	bfs(cur.Right, ans)
}

// @lc code=end
