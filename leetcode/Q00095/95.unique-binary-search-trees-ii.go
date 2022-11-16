/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
 */
package q00095

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
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return bst(1, n)
}

func bst(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	resp := []*TreeNode{}
	for i := start; i <= end; i++ {
		left := bst(start, i-1)
		right := bst(i+1, end)
		for _, l := range left {
			for _, r := range right {
				curr := &TreeNode{i, nil, nil}
				curr.Left = l
				curr.Right = r
				resp = append(resp, curr)
			}
		}
	}
	return resp
}

// @lc code=end
