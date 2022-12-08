/*
 * @lc app=leetcode id=938 lang=golang
 *
 * [938] Range Sum of BST
 */
package q00938

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
func rangeSumBST(root *TreeNode, low int, high int) int {
    if root == nil {
		return 0
	}
	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}
	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}
	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}

func v2(root *TreeNode, l, r int) int {
	sum := 0
	if root == nil {
		return 0
	}
	if l <= root.Val && root.Val <= r {
		sum = root.Val;
	}

	if l <= root.Val || r <= root.Val {
		sum += rangeSumBST(root.Left, l, r);
	} 
	if root.Val <= l || root.Val <= r {
		sum += rangeSumBST(root.Right, l, r);
	}
	return sum    
}
// @lc code=end

