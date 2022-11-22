/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 */
package q00102

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
func levelOrder(root *TreeNode) [][]int {
    resp := [][]int{}

	var traversal func(cur *TreeNode, level int) 
	traversal = func(cur *TreeNode, level int) {
		if cur == nil {
			return
		}
		if len(resp) + 1 == level {
			resp = append(resp, make([]int, 0))
		}
		resp[level-1] = append(resp[level-1], cur.Val)
		traversal(cur.Left, level+1)
		traversal(cur.Right, level+1)
	}

	traversal(root, 1)
	return resp
}



// @lc code=end

