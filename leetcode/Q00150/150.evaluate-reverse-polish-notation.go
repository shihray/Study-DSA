/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 */
package q00150

import "strconv"

// @lc code=start
func evalRPN(tokens []string) int {
    stack := NewStack()

	for _, val := range tokens {

		integer, _ := strconv.Atoi(val)

		switch val {
		case "+", "-", "*", "/":
			val1 := stack.Pop() 
			val2 := stack.Pop()
			var insertData int
			if val == "+" {
				insertData = val2 + val1
			} else if val == "-" {
				insertData = val2 - val1
			} else if val == "*" {
				insertData = val2 * val1
			} else if val == "/" {
				insertData = val2 / val1
			}
			stack.Push(insertData)
		default:
			stack.Push(integer)
		}
	}

	return stack.Pop()
}

type Stack struct {
	Queue []int
}

func NewStack() Stack {
	return Stack{Queue: []int{}}
}
func (s *Stack) Pop() int {
	if len(s.Queue) == 0 {
		return 0
	}
	result := s.Queue[len(s.Queue)-1]
	
	s.Queue = s.Queue[:len(s.Queue)-1]
	return result
}

func (s *Stack) Push(val int) {
	s.Queue = append(s.Queue, val)
}

// @lc code=end

