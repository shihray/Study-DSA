/*
 * @lc app=leetcode id=834 lang=golang
 *
 * [834] Sum of Distances in Tree
 */
package q00834

// @lc code=start
func sumOfDistancesInTree(n int, edges [][]int) []int {
	graph := make([][]int, n)

	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	res := make([]int, n)
	count := make([]int, n)

	var postDfs func(node, parent int)
	postDfs = func(node, parent int) {
		for _, dest := range graph[node] {
			if dest != parent {
				postDfs(dest, node)
				res[node] += res[dest] + count[dest]
				count[node] += count[dest]
			}
		}
		count[node]++
	}

	var preDfs func(node, parent int)
	preDfs = func(node, parent int) {
		for _, dest := range graph[node] {
			if dest != parent {

				res[dest] = res[node] - count[dest] + n - count[dest]
				preDfs(dest, node)
			}
		}
	}

	postDfs(0, -1)
	preDfs(0, -1)
	return res
}

// @lc code=end
