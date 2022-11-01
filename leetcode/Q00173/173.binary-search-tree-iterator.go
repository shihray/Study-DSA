/*
 * @lc app=leetcode id=173 lang=golang
 *
 * [173] Binary Search Tree Iterator
 */
package q00173

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
type BSTIterator struct {
	node []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	st := BSTIterator{}

	for root != nil {
		st.node = append(st.node, root)
		root = root.Left
	}
	return st
}

func (this *BSTIterator) Next() int {
	size := len(this.node)

	cur := this.node[size-1]
	this.node = this.node[:size-1]
	res := cur.Val
	cur = cur.Right
	for cur != nil {
		this.node = append(this.node, cur)
		cur = cur.Left
	}
	return res
}

func (this *BSTIterator) HasNext() bool {
	return len(this.node) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end
