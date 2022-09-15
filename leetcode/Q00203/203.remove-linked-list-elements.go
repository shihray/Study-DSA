/*
 * @lc app=leetcode id=203 lang=golang
 *
 * [203] Remove Linked List Elements
 */
package q00203

type ListNode struct {
    Val int
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
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	curr := head
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else{
			curr = curr.Next
		}
	}
	if head.Val == val {
		head = head.Next
	}
	return head
}
// @lc code=end

