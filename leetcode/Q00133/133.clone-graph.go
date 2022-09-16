/*
 * @lc app=leetcode id=133 lang=golang
 *
 * [133] Clone Graph
 */
package q00133

type Node struct {
    Val int
    Neighbors []*Node
}

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func CloneGraph(node *Node) *Node {

	if node == nil {
		return nil
	}

    recordNode := make(map[*Node]*Node)

	Dfs(node, recordNode)

    return recordNode[node]
}

func Dfs(node *Node, recordNode map[*Node]*Node) {
	if node == nil {
		return
	}
	
	newNode := new(Node)
    newNode.Val = node.Val

	recordNode[node] = newNode

	for _, value := range node.Neighbors {
		if _, ok := recordNode[value]; !ok {
			Dfs(value, recordNode)
		}
		newNode.Neighbors = append(newNode.Neighbors, recordNode[value])
	}
}

// @lc code=end

