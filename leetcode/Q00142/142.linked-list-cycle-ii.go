/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */
package q00142

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
 func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast, isCycle := head, head, false

	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
        if slow == fast {
            isCycle = true
            break
        }
	}
	if !isCycle {
		return nil
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
// @lc code=end

