/*
 * @lc app=leetcode id=875 lang=golang
 *
 * [875] Koko Eating Bananas
 */
package q00875

// @lc code=start
func minEatingSpeed(piles []int, h int) int {
	left, right := 1, max(piles)

	for left < right {
		mid := (left + right) / 2
		if canEat(mid, h, piles) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func canEat(k, h int, piles []int) bool {
	spent := 0

	for _, p := range piles {
		spent += (p / k)
		p = p % k
		if p > 0 {
			spent++
		}
	}
	return spent <= h
}

func max(piles []int) int {
	maxVal := 0
	for _, v := range piles {
		if maxVal < v {
			maxVal = v
		}
	}
	return maxVal
}

// @lc code=end
