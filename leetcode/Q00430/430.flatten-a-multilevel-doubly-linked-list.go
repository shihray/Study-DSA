/*
 * @lc app=leetcode id=430 lang=golang
 *
 * [430] Flatten a Multilevel Doubly Linked List
 */

package q00430

type Node struct {
    Val int
    Prev *Node
    Next *Node
    Child *Node
}
// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
    curr := root

	for curr != nil {
		if curr.Child == nil {
			curr = curr.Next
		} else {
			next := curr.Next
			node := flatten(curr.Child)
			curr.Child = nil
			curr.Next, node.Prev = node, curr
			for curr.Next != nil {
				curr = curr.Next
			}
			if next == nil {
				break
			}
			curr.Next, next.Prev = next, curr
			curr = curr.Next
		}
	}
	return root
}
// @lc code=end

