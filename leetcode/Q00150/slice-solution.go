/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 */
package q00150

import "strconv"

// @lc code=start
func evalRPNSlice(tokens []string) int {

	var stack []int
    for i := 0; i < len(tokens); i++ {
		
		cur := tokens[i]
		switch cur {
		case "+", "-", "*", "/":
			a, b := stack[len(stack)-2] , stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, getVal(a, b, cur))
		default:
			val, _ := strconv.Atoi(cur)
			stack = append(stack, val)
		}
	}

	return stack[0]
}

func getVal(a, b int, s string) int {
	switch s {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	return 0
}
// @lc code=end

