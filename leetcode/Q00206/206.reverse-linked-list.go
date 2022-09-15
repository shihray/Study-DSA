/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */
package q00206

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
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
    for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}
// @lc code=end

