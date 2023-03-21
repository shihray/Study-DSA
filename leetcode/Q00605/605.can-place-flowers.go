/*
 * @lc app=leetcode id=605 lang=golang
 *
 * [605] Can Place Flowers
 */
package q00605

// @lc code=start
func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		curr := flowerbed[i]
		if curr == 0 {
			prev := 0
			next := 0
			// set edge prev and next to 0
			if i > 0 {
				prev = flowerbed[i-1]
			}
			if i < len(flowerbed)-1 {
				next = flowerbed[i+1]
			}

			// if no flower in prev and next, put the flower
			if prev == 0 && next == 0 {
				flowerbed[i] = 1 // place the flower
				n--
			}
		}
		if n <= 0 {
			return true
		}
	}
	return false
}

// @lc code=end
