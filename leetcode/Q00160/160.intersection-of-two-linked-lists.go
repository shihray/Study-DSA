/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 */

package leetcode

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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
		return nil
	}

	listA, listB := headA, headB
	for listA != listB {
		listA, listB = listA.Next, listB.Next
		if listA == listB {
			return listA
		}
		if listA == nil {
			listA = headB
		}
		if listB == nil {
			listB = headA
		}
	}
	return listA
}
// @lc code=end

