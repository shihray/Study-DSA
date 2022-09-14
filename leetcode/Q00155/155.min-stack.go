/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */
package leetcode

// @lc code=start
type MinStackStuct struct {
	Val int
	Min int
}

type MinStack struct {
    queue []MinStackStuct
}


func Constructor() MinStack {
    return MinStack {
		queue: []MinStackStuct{},
	}
}


func (this *MinStack) Push(val int)  {
	if len(this.queue) > 0 {
		this.queue = append(this.queue, MinStackStuct{
			Val: val,
			Min: min(val, this.queue[len(this.queue)-1].Min),
		})
	} else {
		this.queue = append(this.queue, MinStackStuct{
			Val: val,
			Min: val,
		})
	}
    
}


func (this *MinStack) Pop()  {
	if len(this.queue) == 0 {
		return
	}
	this.queue = this.queue[:len(this.queue)-1]
}


func (this *MinStack) Top() int {
    return this.queue[len(this.queue)-1].Val
}


func (this *MinStack) GetMin() int {
    
    return this.queue[len(this.queue)-1].Min
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

