/*
 * @lc app=leetcode id=1443 lang=golang
 *
 * [1443] Minimum Time to Collect All Apples in a Tree
 */
package q01443

// @lc code=start
func minTime(n int, edges [][]int, hasApple []bool) int {
	graph, visit := make(map[int][]int), make(map[int]bool)

	for _, v := range edges {
		if _, ok := graph[v[0]]; !ok {
			graph[v[0]] = []int{}
		}
		graph[v[0]] = append(graph[v[0]], v[1])
		if _, ok := graph[v[1]]; !ok {
			graph[v[1]] = []int{}
		}
		graph[v[1]] = append(graph[v[1]], v[0])
	}
	visit[0] = true

	return dfs(0, graph, visit, hasApple)
}

func dfs(cur int, graph map[int][]int, visit map[int]bool, hasApple []bool) int {
	visit[cur] = true
	resp := 0

	for _, v := range graph[cur] {
		if !visit[v] {
			resp += dfs(v, graph, visit, hasApple)
		}
	}

	if cur == 0 {
		return resp
	}
	if resp > 0 || hasApple[cur] {
		resp += 2
	}
	return resp
}

// @lc code=end
