/*
 * @lc app=leetcode id=1971 lang=golang
 *
 * [1971] Find if Path Exists in Graph
 */
package q01971

// @lc code=start
func validPath(n int, edges [][]int, source int, destination int) bool {
    e := make(map[int][]int)

	for _, v := range edges {
		e[v[0]] = append(e[v[0]], v[1])
		e[v[1]] = append(e[v[1]], v[0])
	}

	seen := make(map[int]bool)
	stack := []int{source}

	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if cur == destination {
			return true
		}

		if !seen[cur] {
			stack = append(stack, e[cur]...)
			seen[cur] = true
		}
	}
	return false
}
// @lc code=end

