/*
 * @lc app=leetcode id=1519 lang=golang
 *
 * [1519] Number of Nodes in the Sub-Tree With the Same Label
 */
package q01519

// @lc code=start
func countSubTrees(n int, edges [][]int, labels string) []int {
	var labelCounter [26]int

	graph := make([][]int, n)
	resp := make([]int, n)

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	var dfs func(cur int, parent int)
	dfs = func(cur int, parent int) {
		nodeId := labels[cur] - 97

		per := labelCounter[nodeId]
		labelCounter[nodeId]++

		for _, v := range graph[cur] {
			if v == parent {
				continue
			}
			dfs(v, cur)
		}
		resp[cur] = labelCounter[nodeId] - per
	}

	dfs(0, -1)
	return resp
}

// @lc code=end
