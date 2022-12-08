/*
 * @lc app=leetcode id=872 lang=golang
 *
 * [872] Leaf-Similar Trees
 */
package q00872

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
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    var arr1, arr2 []int

	findLeaf(root1, &arr1)
	findLeaf(root2, &arr2)

	if len(arr1) != len(arr2) {
		return false
	}

	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func findLeaf(root *TreeNode, arr *[]int) {
	if root == nil {
		return 
	}
	if root.Left == nil && root.Right == nil {
		*arr = append(*arr, root.Val)
	}
	findLeaf(root.Left, arr)
	findLeaf(root.Right, arr)
}
// @lc code=end

