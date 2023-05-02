/*
 * @lc app=leetcode id=1822 lang=golang
 *
 * [1822] Sign of the Product of an Array
 */
package q01822

// @lc code=start
func arraySign(nums []int) int {
	sign := 1

    for _, num := range nums {
        if num <= -1 {
            sign = - sign
        } else if num == 0 {
            return 0
        }
    }

    return sign
}

// @lc code=end
