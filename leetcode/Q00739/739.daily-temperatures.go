/*
 * @lc app=leetcode id=739 lang=golang
 *
 * [739] Daily Temperatures
 */
package q00739

// @lc code=start
type TempSt struct {
	Temp int
	Index int
}

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))

	var stack []TempSt

	for i := len(temperatures)-1; i >= 0; i-- {
		
		result[i] = 0
		temp := temperatures[i]

		for len(stack) > 0 && (temp >= stack[len(stack)-1].Temp) {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			result[i] = stack[len(stack) - 1].Index - i
		}

		stack = append(stack, TempSt{Temp: temp, Index: i})
	}


	return result
}
// @lc code=end

