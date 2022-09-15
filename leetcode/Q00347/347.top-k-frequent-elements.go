/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */
package q00347

import "container/heap"

// @lc code=start
func topKFrequent(nums []int, k int) []int {
    m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}

	heapData := make(MyNums, len(m))
	i:=0
	for k, v := range m {
		heapData[i] = MyHeap{
			Val: k,
			Count: v,
		}
		i++
	}
	heap.Init(&heapData)
	var res []int
	for i := 0; i < k; i++ {
		num := heap.Pop(&heapData).(MyHeap)
		res = append(res, num.Val)
	}
	return res
}


type MyHeap struct {
	Val int
	Count int
} 

type MyNums []MyHeap

func (n MyNums) Len() int {
	return len(n)
}

func (n MyNums) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n MyNums) Less(i, j int) bool {
	return n[i].Count >= n[j].Count
}

func (n *MyNums) Push(num interface{}) {
	myNum := num.(MyHeap)
	*n = append(*n, myNum)
}

func (n *MyNums) Pop() interface{} {
	tmp := *n
	l := len(tmp)
	var res interface{} = tmp[l-1]
	*n = tmp[:l-1]
	return res
}


// @lc code=end

