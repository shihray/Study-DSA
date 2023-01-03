/*
 * @lc app=leetcode id=1834 lang=golang
 *
 * [1834] Single-Threaded CPU
 */
package q01834

import (
	"container/heap"
	"sort"
)

// @lc code=start
func getOrder(tasks [][]int) []int {

	list := make(PQueue, len(tasks))
	for i := 0; i < len(tasks); i++ {
		list[i] = &Item{
			val:      tasks[i][0],
			priority: tasks[i][1],
			idx:      i,
		}
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].val < list[j].val
	})

	var (
		resp []int
		cur  = list[0].val
		i    = 0
		pq   = PQueue{}
	)

	for pq.Len() > 0 || i < list.Len() {
		for i < len(list) && list[i].val <= cur {
			heap.Push(&pq, list[i])
			i++
		}
		if pq.Len() == 0 {
			cur = list[i].val
		} else {
			tmp := heap.Pop(&pq).(*Item)
			cur += tmp.priority
			resp = append(resp, tmp.idx)
		}

	}
	return resp
}

type Item struct {
	val      int
	priority int
	idx      int
}

type PQueue []*Item

func (pq PQueue) Len() int { return len(pq) }

func (pq PQueue) Less(i, j int) bool {
	if pq[i].priority == pq[j].priority {
		return pq[i].idx < pq[j].idx
	}
	return pq[i].priority < pq[j].priority
}

func (pq PQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	*pq = old[:n-1]
	return old[n-1]
}

// @lc code=end
