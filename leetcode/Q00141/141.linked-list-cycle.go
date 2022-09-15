/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

package q00141

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
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast.Next != nil && fast.Next.Next != nil && fast != slow{
        slow = slow.Next
        fast = fast.Next.Next
    }
    
    return slow == fast
}
// @lc code=end

