/*
 * @lc app=leetcode id=1319 lang=golang
 *
 * [1319] Number of Operations to Make Network Connected
 */
package q01319

// @lc code=start
func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}

	father := make([]int, n)

	for i := range father {
		father[i] = i
	}

	for _, c := range connections {
		union(c[0], c[1], father)
	}

	anc := map[int]bool{}

	for i := range father {
		anc[find(i, father)] = true
	}
	return len(anc) - 1
}

func union(i, j int, father []int) {
	fi := find(i, father)
	fj := find(j, father)

	if fi < fj {
		father[fj] = fi
	} else if fi > fj {
		father[fi] = fj
	}
}

func find(i int, father []int) int {
	for father[i] != i {
		i, father[i] = father[i], father[father[i]]
	}
	return i
}

// @lc code=end
