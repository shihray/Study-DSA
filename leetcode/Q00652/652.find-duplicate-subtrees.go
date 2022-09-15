/*
 * @lc app=leetcode id=652 lang=golang
 *
 * [652] Find Duplicate Subtrees
 */
package q00652

import "fmt"

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
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    hashMap := map[string]int{}
	duplicate := []*TreeNode{}

	dfs(root, hashMap, &duplicate)
	return duplicate
}

func dfs(node *TreeNode, hashMap map[string]int, duplicate *[]*TreeNode) string {
	if node == nil {
		return "nil"
	}

	lStr := dfs(node.Left, hashMap, duplicate)
	rStr := dfs(node.Right, hashMap, duplicate)
	buildStr := fmt.Sprintf("(%s),(%v),(%s)", lStr, node.Val, rStr)

	hashMap[buildStr]++

	if hashMap[buildStr] == 2 {
		*duplicate = append(*duplicate, node)
	}
	return buildStr
}
// @lc code=end

