/*
 * @lc app=leetcode id=1339 lang=golang
 *
 * [1339] Maximum Product of Splitted Binary Tree
 */
package q01339

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
func maxProduct(root *TreeNode) int {

    pos := make([]int64, 0)
	mod := int64(1e9) + 7
	resp := int64(0)

	totalSum := dfs(root, &pos)

	for _, v := range pos {
		cur := int64(totalSum - v) * int64(v)
		resp = max(resp, cur)
	}

	return int(resp % mod)
}

func dfs(root *TreeNode, pos *[]int64) int64 {
	if root == nil {
		return 0
	}
	leftSum := dfs(root.Left, pos)
	rightSum := dfs(root.Right, pos)

	tmp := leftSum + rightSum + int64(root.Val)
	*pos = append(*pos, tmp)
	return tmp
}

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}
// @lc code=end

