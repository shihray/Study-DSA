/*
 * @lc app=leetcode id=2343 lang=golang
 *
 * [2343] Query Kth Smallest Trimmed Number
 */
package q02343

import (
	"fmt"
	"sort"
	"strconv"
)

// @lc code=start
func smallestTrimmedNumbers(nums []string, queries [][]int) []int {

	result := make([]int, len(queries))

	for i, query := range queries {
		
		var currentNums [][2]string
		k, trim := query[0], query[1]

		for numIdx, num := range nums {
			currentNums = append(currentNums, [2]string{num[len(num)-trim:], fmt.Sprintf("%v", numIdx)})
		}

		sort.Slice(currentNums, func(i, j int) bool {
			if currentNums[i][0] != currentNums[j][0] {
				return isSmaller(currentNums[i][0], currentNums[j][0])
			} else {
                return isSmaller(currentNums[i][1], currentNums[j][1])
			}
		})

		result[i] = toInt(currentNums[k-1][1])
	}
	return result
}

func toInt(s string) int {
	res, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return res
}

func isSmaller(s1 string, s2 string) bool {
	if len(s1) < len(s2) {
		return true
	}
	
	if len(s1) > len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] < s2[i] {
			return true
		} else if s1[i] > s2[i] {
			return false
		}
	}
	return true
}
// @lc code=end

