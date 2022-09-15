/*
 * @lc app=leetcode id=328 lang=golang
 *
 * [328] Odd Even Linked List
 */

package q00328

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
func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}
	count := 1
	var oddHead, oddTail *ListNode
    var evenHead, evenTail *ListNode

	for head != nil {
		if count % 2 == 1 {
			if oddHead == nil {
				oddHead = head
				oddTail = oddHead
			} else {
				oddTail.Next = head
                oddTail = oddTail.Next
			}
		} else {
			if evenHead == nil {
				evenHead = head
				evenTail = evenHead
			} else {
				evenTail.Next = head
                evenTail = evenTail.Next
			}
		}
		count++
		head, head.Next = head.Next, nil
	}

	oddTail.Next = evenHead

	return oddHead
}
// @lc code=end

