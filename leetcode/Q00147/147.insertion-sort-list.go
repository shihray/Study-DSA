/*
 * @lc app=leetcode id=147 lang=golang
 *
 * [147] Insertion Sort List
 */
package q00147

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
func insertionSortList(head *ListNode) *ListNode {
    
	newList := new(ListNode)

	for head != nil {
		cur := newList
		for ;cur.Next != nil && cur.Next.Val < head.Val;  cur = cur.Next {
		}
		cur.Next, head.Next, head = head, cur.Next, head.Next
	}
	return newList.Next
}
// @lc code=end
