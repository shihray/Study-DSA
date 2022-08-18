/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 */
package leetcode

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    
	curr, end := head, head
	length := 1

	for end != nil && end.Next != nil {
		end = end.Next
		length++
	}


	for i := 0; i < length - (k % length); i++ {
		if end != nil {
			end.Next = &ListNode{Val: curr.Val}
			end = end.Next
			curr = curr.Next
		}
	}
	return curr
}
// @lc code=end

