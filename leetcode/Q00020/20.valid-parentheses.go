/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */
package q00020

// @lc code=start
func isValid(s string) bool {
    if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	runeMap := map[rune]rune{
		'(' : ')',
		'[' : ']',
		'{' : '}',
	}

	stack := []rune{}
	for _, r := range s {
		if _, ok := runeMap[r]; ok {
			stack = append(stack, r)
		} else if len(stack) == 0 || runeMap[stack[len(stack)-1]] != r {
			return false
		} else {
			stack = stack[: len(stack)-1]
		}
	}
	return len(stack) == 0
}


// @lc code=end

