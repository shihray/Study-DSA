/*
 * @lc app=leetcode id=382 lang=golang
 *
 * [382] Linked List Random Node
 */
package q00382

import (
	"math/rand"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {

	return Solution{
		head: head,
	}
}

func (this *Solution) GetRandom() int {
	m := float64(1)
	mv := 0
	for n := this.head; n != nil; n = n.Next {
		r := rand.Float64()
		if r < m {
			m = r
			mv = n.Val
		}
	}
	return mv
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
// @lc code=end
