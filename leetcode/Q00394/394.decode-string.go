/*
 * @lc app=leetcode id=394 lang=golang
 *
 * [394] Decode String
 */
package q00394

import (
	"strconv"
	"strings"
)

// @lc code=start
func decodeString(s string) string {
    intStack := MyStack{}
	strStack := MyStack{}

	k := 0
	curStr := ""
	for i := 0; i < len(s); i++ {
		
		chr := string(s[i])
		if n, err := strconv.Atoi(chr); err == nil {
			k = k * 10 + n
		} else if chr == "[" {
			intStack.Push(k)
			strStack.Push(curStr)
			k = 0
			curStr = ""
		} else if chr == "]" {
			tmp := curStr
			
			intVal := intStack.Pop().(int)
			curStr = strStack.Pop().(string)

			curStr += strings.Repeat(tmp, intVal)
		} else {
			curStr += chr
		}
	}
	return curStr
}

type MyStack struct {
    stack []interface{}
}

func (s *MyStack) Pop() interface{} {
	var result interface{}
	if len(s.stack) == 0 {
		return result
	}
	result = s.stack[len(s.stack) - 1]
	s.stack = s.stack[:len(s.stack)-1]
	return result
}

func (s *MyStack) Push(x interface{}) {
	s.stack = append(s.stack, x)
}

// @lc code=end

