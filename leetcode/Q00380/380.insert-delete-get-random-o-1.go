/*
 * @lc app=leetcode id=380 lang=golang
 *
 * [380] Insert Delete GetRandom O(1)
 */
package q00380

import "math/rand"

// @lc code=start
type RandomizedSet struct {
    Item map[int]bool
	CurrArr []int
}

func Constructor() RandomizedSet {
    return RandomizedSet{
		Item: map[int]bool{},
		CurrArr: make([]int, 0),
	}
}


func (this *RandomizedSet) Insert(val int) bool {
    if item, ok := this.Item[val]; ok && item {
		return false
	} 
	this.Item[val] = true
	this.CurrArr = append(this.CurrArr, val)
	return true
}


func (this *RandomizedSet) Remove(val int) bool {
    if item, ok := this.Item[val]; ok {
		this.Item[val] = false

		for i := 0; i < len(this.CurrArr); i++ {
			if this.CurrArr[i] == val {
				this.CurrArr = append(this.CurrArr[:i], this.CurrArr[i+1:]...)
			}
		}

		return item
	}
	return false
}


func (this *RandomizedSet) GetRandom() int {
	i := rand.Intn(len(this.CurrArr))
	return this.CurrArr[i]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end

