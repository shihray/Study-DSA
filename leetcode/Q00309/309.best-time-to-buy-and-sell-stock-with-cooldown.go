/*
 * @lc app=leetcode id=309 lang=golang
 *
 * [309] Best Time to Buy and Sell Stock with Cooldown
 */
package q00309

// @lc code=start
func maxProfit(prices []int) int {

	if len(prices) <= 1 {
		return 0
	}

	stack := make([]int, len(prices)+1)
	min_price := prices[0]

	for i := 1; i <= len(prices); i++ {
		if i >= 2 {
			min_price = min(prices[i-1]-stack[i-2], min_price)
		}
		stack[i] = max(stack[i-1], prices[i-1]-min_price)
	}
	return stack[len(prices)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
