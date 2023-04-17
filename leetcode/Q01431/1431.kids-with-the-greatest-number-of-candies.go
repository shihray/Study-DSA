/*
 * @lc app=leetcode id=1431 lang=golang
 *
 * [1431] Kids With the Greatest Number of Candies
 */
package q01431

// @lc code=start
func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCnt := getMax(candies)
	n := len(candies)

	res := make([]bool, n)

	for idx, candy := range candies {
		if candy + extraCandies >= maxCnt {
			res[idx] = true
		}
	}
	return res
}

func getMax(arr []int) int {
	res := 0
	for _, v := range arr {
		if res < v {
			res = v
		}
	}
	return res
}

// @lc code=end
