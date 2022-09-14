/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
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
func isPalindrome(head *ListNode) bool {
    
	if head == nil {
		return false
	}

	var arr []int

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	return palindrome(arr)
}

func palindrome(arr []int) bool {

	for i, j := 0, len(arr)-1; i < len(arr); i, j = i+1,j-1 {
		if arr[i] != arr[j] {
			return false
		} 
	}
	return true
}
// @lc code=end

