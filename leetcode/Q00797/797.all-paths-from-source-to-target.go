/*
 * @lc app=leetcode id=797 lang=golang
 *
 * [797] All Paths From Source to Target
 */
package q00797

// @lc code=start
func allPathsSourceTarget(graph [][]int) [][]int {

	if len(graph) == 0 {
		return [][]int{}
	}
	resp := make([][]int, 0)
	tmp := make([]int, 0)

	Dfs(0, graph, &resp, &tmp)
	return resp
}

func Dfs(curr int, graph [][]int, resp *[][]int, temp *[]int) {

	(*temp) = append(*temp, curr)
	if curr == len(graph)-1 {
		np := make([]int, len(*temp))
		copy(np, *temp)
		*resp = append(*resp, np)
	} else {
		for _, v := range graph[curr] {
			Dfs(v, graph, resp, temp)
			(*temp) = (*temp)[:len(*temp)-1]
		}
	}
}

// @lc code=end
