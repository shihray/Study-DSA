/*
 * @lc app=leetcode id=904 lang=golang
 *
 * [904] Fruit Into Baskets
 */
package q00904

// @lc code=start
func totalFruit(f []int) int {

	maxPick := 0
	basket := map[int]int{}
	left := 0

	for right := 0; right < len(f); right++ {

		basket[f[right]]++

		for len(basket) > 2 {
			basket[f[left]]--
			if basket[f[left]] == 0 {
				delete(basket, f[left])
			}
			left++
		}
		maxPick = max(maxPick, right-left+1)
	}
	return maxPick
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
