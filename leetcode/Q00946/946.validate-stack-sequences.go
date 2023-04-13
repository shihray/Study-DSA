/*
 * @lc app=leetcode id=946 lang=golang
 *
 * [946] Validate Stack Sequences
 */
package q00946

// @lc code=start
func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	count := 0 

	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) != 0 && popped[count] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			count++
			if count == len(popped) {
				break
			}
		}
	}

	return len(stack) == 0
}

// @lc code=end
