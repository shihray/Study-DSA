/*
 * @lc app=leetcode id=983 lang=golang
 *
 * [983] Minimum Cost For Tickets
 */
package q00983

// @lc code=start
func mincostTickets(days []int, costs []int) int {
	stack7, stack30 := make([]*node, 0, len(days)), make([]*node, 0, len(days))
	cost := 0
	for _, day := range days {
		for len(stack7) > 0 && stack7[0].day+7 <= day {
			stack7 = stack7[1:]
		}
		for len(stack30) > 0 && stack30[0].day+30 <= day {
			stack30 = stack30[1:]
		}
		stack7 = append(stack7, &node{day, cost + costs[1]})
		stack30 = append(stack30, &node{day, cost + costs[2]})
		cost = min(cost+costs[0], stack7[0].cost, stack30[0].cost)
	}
	return cost
}

type node struct {
	day  int
	cost int
}

func min(nums ...int) int {
	m := nums[0]
	for _, num := range nums {
		if num < m {
			m = num
		}
	}
	return m
}

// @lc code=end
