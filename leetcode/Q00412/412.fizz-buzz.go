/*
 * @lc app=leetcode id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

package q00412

import "strconv"

// @lc code=start
func fizzBuzz(n int) []string {

	res := make([]string, n)
    for i := 1; i <= n; i++ {
		curStr := strconv.Itoa(i)
		if i%3 == 0 && i%5 == 0 {
			curStr = "FizzBuzz"
		} else if i%3 == 0 {
			curStr = "Fizz"
		} else if i%5 == 0 {
			curStr = "Buzz"
		}
		res[i-1] = curStr
	}
	return res
}
// @lc code=end

