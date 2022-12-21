/*
 * @lc app=leetcode id=886 lang=golang
 *
 * [886] Possible Bipartition
 */
package q00886

// @lc code=start
func possibleBipartition(n int, dislikes [][]int) bool {
	graph := make([][]int, n+1)

	for _, v := range dislikes {
		a, b := v[0] - 1, v[1] - 1
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	colors := make([]int, n+1)
	for i := 0; i < n; i++ {
		if colors[i] == 0 && !isValid(&graph, &colors, i, 1) {
			return false
		}
	}
	return true
}

func isValid(graph *[][]int, colors *[]int, i, color int) bool {
	if (*colors)[i] != 0 {
		return (*colors)[i] == color
	}

	(*colors)[i] = color
	for _, j := range (*graph)[i] {
		if (*colors)[i] == (*colors)[j] || !isValid(graph, colors, j, -color) {
			return false
		}
	}
	return true
}
// @lc code=end

