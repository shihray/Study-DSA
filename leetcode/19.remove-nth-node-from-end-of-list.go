/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	
	slow, fast := dummy, dummy
	count := 0
	for fast.Next != nil {
		fast = fast.Next

		count++
		if count > n {
			slow = slow.Next
		}
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
// @lc code=end

