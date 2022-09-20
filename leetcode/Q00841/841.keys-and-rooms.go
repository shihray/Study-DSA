/*
 * @lc app=leetcode id=841 lang=golang
 *
 * [841] Keys and Rooms
 */
package q00841

// @lc code=start
func canVisitAllRooms(rooms [][]int) bool {

	m := map[int]bool{}
	Dfs(0, m, rooms)

	return len(m) == len(rooms)
}

func Dfs(cur int, m map[int]bool, rooms [][]int) {
	if _, ok := m[cur]; ok {
		return
	}
	m[cur] = true

	for _, next := range rooms[cur] {
		Dfs(next, m, rooms)
	}
}
// @lc code=end

